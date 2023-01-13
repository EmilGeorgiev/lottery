package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/EmilGeorgiev/lottery/testutil/keeper"
	"github.com/EmilGeorgiev/lottery/testutil/nullify"
	"github.com/EmilGeorgiev/lottery/x/lottery/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFinishedLottery(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FinishedLottery {
	items := make([]types.FinishedLottery, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFinishedLottery(ctx, items[i])
	}
	return items
}

func TestFinishedLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNFinishedLottery(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFinishedLottery(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFinishedLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNFinishedLottery(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFinishedLottery(ctx,
			item.Index,
		)
		_, found := keeper.GetFinishedLottery(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFinishedLotteryGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNFinishedLottery(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFinishedLottery(ctx)),
	)
}
