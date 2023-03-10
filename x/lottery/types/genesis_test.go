package types_test

import (
	"testing"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				Lottery: types.Lottery{
					EnterLotteryTxs: []*types.EnterLotteryTx{
						{UserAddress: "address1", Bet: 1},
						{UserAddress: "address2", Bet: 2},
						{UserAddress: "address3", Bet: 3},
					},
				},
				SystemInfo: types.SystemInfo{
					NextId:      6,
					LotteryPool: 63,
				},
				FinishedLotteryList: []types.FinishedLottery{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated finishedLottery",
			genState: &types.GenesisState{
				FinishedLotteryList: []types.FinishedLottery{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDefaultGenesisState_ExpectedInitialNextId(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			Lottery:             types.Lottery{},
			SystemInfo:          types.SystemInfo{NextId: 1, LotteryPool: 0},
			FinishedLotteryList: []types.FinishedLottery{},
		},
		types.DefaultGenesis())
}
