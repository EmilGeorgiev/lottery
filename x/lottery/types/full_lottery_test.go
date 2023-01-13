package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

const alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"

func TestLottery_RegisterNewUser(t *testing.T) {
	// SetUp
	cases := []struct {
		name         string
		enterLottery *types.MsgEnterLottery
		current      types.Lottery
		expect       types.Lottery
	}{
		{
			name:         "The first user emter the lottery",
			enterLottery: &types.MsgEnterLottery{Creator: "address1", Bet: 1},
			current:      types.Lottery{},
			expect: types.Lottery{Users: []*types.User{
				{Address: "address1", Bet: 1},
			}},
		},
		{
			name:         "New user enter the lottery",
			enterLottery: &types.MsgEnterLottery{Creator: "address2", Bet: 2},
			current: types.Lottery{Users: []*types.User{
				{Address: "address1", Bet: 1},
			}},
			expect: types.Lottery{Users: []*types.User{
				{Address: "address1", Bet: 1},
				{Address: "address2", Bet: 2},
			}},
		},
		{
			name:         "User entered lottery more then once",
			enterLottery: &types.MsgEnterLottery{Creator: "address2", Bet: 5},
			current: types.Lottery{Users: []*types.User{
				{Address: "address1", Bet: 1},
				{Address: "address2", Bet: 2},
				{Address: "address3", Bet: 3},
			}},
			expect: types.Lottery{Users: []*types.User{
				{Address: "address1", Bet: 1},
				{Address: "address2", Bet: 5},
				{Address: "address3", Bet: 3},
			}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Action
			c.current.RegisterNewUser(c.enterLottery)

			// Assert
			require.Equal(t, c.expect, c.current)
		})
	}
}

func TestMsgEnteredLottery_GetAddress(t *testing.T) {
	// SetUp
	el := types.MsgEnterLottery{
		Creator: alice,
		Bet:     3,
		Denom:   "token",
	}

	// Action
	addr, err := el.GetAddress()

	// Assert
	expected, err1 := sdk.AccAddressFromBech32(alice)
	require.NoError(t, err)
	require.NoError(t, err1)
	require.Equal(t, expected, addr)
}

func TestMsgEnteredLottery_GetAddressFailed(t *testing.T) {
	// SetUp
	el := types.MsgEnterLottery{
		Creator: "invalid",
		Bet:     3,
		Denom:   "token",
	}

	// Action
	addr, err := el.GetAddress()

	// Assert
	require.Error(t, err)
	require.Nil(t, addr)
}
