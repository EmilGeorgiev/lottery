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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Lottery, got.Lottery)
	// this line is used by starport scaffolding # genesis/test/assert
}
