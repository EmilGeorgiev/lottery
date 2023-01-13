package lottery_test

import (
	"testing"

	keepertest "github.com/EmilGeorgiev/lottery/testutil/keeper"
	"github.com/EmilGeorgiev/lottery/testutil/nullify"
	"github.com/EmilGeorgiev/lottery/x/lottery"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Lottery: types.Lottery{
			Users: []*types.User{
				{Address: "address1", Bet: 10},
				{Address: "address2", Bet: 20},
				{Address: "address3", Bet: 30},
			},
		},
		SystemInfo: types.SystemInfo{
			NextId:      63,
			LotteryPool: 42,
		},
		FinishedLotteryList: []types.FinishedLottery{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Lottery, got.Lottery)
	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.FinishedLotteryList, got.FinishedLotteryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
