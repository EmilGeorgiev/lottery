package types_test

import (
	"testing"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestLottery_RegisterNewUser(t *testing.T) {
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
			c.current.RegisterNewUser(c.enterLottery)
			require.Equal(t, c.expect, c.current)
		})
	}
}
