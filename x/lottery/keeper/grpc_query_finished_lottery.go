package keeper

import (
	"context"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FinishedLotteryAll(c context.Context, req *types.QueryAllFinishedLotteryRequest) (*types.QueryAllFinishedLotteryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var finishedLotterys []types.FinishedLottery
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	finishedLotteryStore := prefix.NewStore(store, types.KeyPrefix(types.FinishedLotteryKeyPrefix))

	pageRes, err := query.Paginate(finishedLotteryStore, req.Pagination, func(key []byte, value []byte) error {
		var finishedLottery types.FinishedLottery
		if err := k.cdc.Unmarshal(value, &finishedLottery); err != nil {
			return err
		}

		finishedLotterys = append(finishedLotterys, finishedLottery)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFinishedLotteryResponse{FinishedLottery: finishedLotterys, Pagination: pageRes}, nil
}

func (k Keeper) FinishedLottery(c context.Context, req *types.QueryGetFinishedLotteryRequest) (*types.QueryGetFinishedLotteryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFinishedLottery(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFinishedLotteryResponse{FinishedLottery: val}, nil
}
