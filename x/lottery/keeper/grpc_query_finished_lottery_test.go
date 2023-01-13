package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/EmilGeorgiev/lottery/testutil/keeper"
	"github.com/EmilGeorgiev/lottery/testutil/nullify"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFinishedLotteryQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFinishedLottery(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetFinishedLotteryRequest
		response *types.QueryGetFinishedLotteryResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetFinishedLotteryRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetFinishedLotteryResponse{FinishedLottery: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetFinishedLotteryRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetFinishedLotteryResponse{FinishedLottery: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetFinishedLotteryRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.FinishedLottery(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestFinishedLotteryQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFinishedLottery(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllFinishedLotteryRequest {
		return &types.QueryAllFinishedLotteryRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FinishedLotteryAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FinishedLottery), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FinishedLottery),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FinishedLotteryAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.FinishedLottery), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.FinishedLottery),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.FinishedLotteryAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.FinishedLottery),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.FinishedLotteryAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
