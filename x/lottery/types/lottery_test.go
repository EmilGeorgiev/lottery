package types_test

import (
	"testing"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
)

func TestLottery_RegisterEnterLotteryTx(t *testing.T) {
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
			expect: types.Lottery{EnterLotteryTxs: []*types.EnterLotteryTx{
				{UserAddress: "address1", Bet: 1},
			}},
		},
		{
			name:         "New user enter the lottery",
			enterLottery: &types.MsgEnterLottery{Creator: "address2", Bet: 2},
			current: types.Lottery{EnterLotteryTxs: []*types.EnterLotteryTx{
				{UserAddress: "address1", Bet: 1},
			}},
			expect: types.Lottery{EnterLotteryTxs: []*types.EnterLotteryTx{
				{UserAddress: "address1", Bet: 1},
				{UserAddress: "address2", Bet: 2},
			}},
		},
		{
			name:         "User entered lottery more then once",
			enterLottery: &types.MsgEnterLottery{Creator: "address2", Bet: 5},
			current: types.Lottery{EnterLotteryTxs: []*types.EnterLotteryTx{
				{UserAddress: "address1", Bet: 1},
				{UserAddress: "address2", Bet: 2},
				{UserAddress: "address3", Bet: 3},
			}},
			expect: types.Lottery{EnterLotteryTxs: []*types.EnterLotteryTx{
				{UserAddress: "address1", Bet: 1},
				{UserAddress: "address2", Bet: 5},
				{UserAddress: "address3", Bet: 3},
			}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Action
			c.current.RegisterNewTx(c.enterLottery)

			for i, tx := range c.current.EnterLotteryTxs {
				c.expect.EnterLotteryTxs[i].Datetime = tx.Datetime
			}

			// Assert
			require.Equal(t, c.expect, c.current)
		})
	}
}

func TestLottery_GetLowestAndHighestBet(t *testing.T) {
	// SetUp
	cases := []struct {
		name            string
		lottery         types.Lottery
		expectedLowest  uint64
		expectedHighest uint64
	}{
		{
			name: "One lowest and one highest values",
			lottery: types.Lottery{
				EnterLotteryTxs: []*types.EnterLotteryTx{
					{Bet: 1}, {Bet: 2}, {Bet: 3}, {Bet: 4}, {Bet: 5}, {Bet: 6},
				},
			},
			expectedHighest: 6,
			expectedLowest:  1,
		},
		{
			name: "More then one lowest and highest values",
			lottery: types.Lottery{
				EnterLotteryTxs: []*types.EnterLotteryTx{
					{Bet: 40}, {Bet: 11}, {Bet: 60}, {Bet: 11}, {Bet: 30}, {Bet: 60},
				},
			},
			expectedHighest: 60,
			expectedLowest:  11,
		},
		{
			name: "All bets are equal",
			lottery: types.Lottery{
				EnterLotteryTxs: []*types.EnterLotteryTx{
					{Bet: 5}, {Bet: 5}, {Bet: 5}, {Bet: 5}, {Bet: 5}, {Bet: 5},
				},
			},
			expectedHighest: 5,
			expectedLowest:  5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Action
			actualLowest, actualHighest := c.lottery.GetLowestAndHighestBet()

			// Assert
			require.Equal(t, c.expectedLowest, actualLowest)
			require.Equal(t, c.expectedHighest, actualHighest)
		})
	}
}

func TestLottery_GetSumOfAllBets(t *testing.T) {
	// SetUp
	cases := []struct {
		name     string
		lottery  types.Lottery
		expected uint64
	}{
		{
			name: "Get sum of all bets",
			lottery: types.Lottery{
				EnterLotteryTxs: []*types.EnterLotteryTx{
					{Bet: 1}, {Bet: 2}, {Bet: 3}, {Bet: 4}, {Bet: 5}, {Bet: 6},
				},
			},
			expected: 21,
		},
		{
			name: "Lottery has only one user",
			lottery: types.Lottery{
				EnterLotteryTxs: []*types.EnterLotteryTx{{Bet: 3}},
			},
			expected: 3,
		},
		{
			name:     "Lottery without users",
			lottery:  types.Lottery{},
			expected: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Action
			actual := c.lottery.GetSumOfAllBets()

			// Assert
			require.Equal(t, c.expected, actual)
		})
	}
}

func TestLottery_GetSumOfAllBetsPlusFees(t *testing.T) {
	// SetUp
	cases := []struct {
		name     string
		lottery  types.Lottery
		expected uint64
	}{
		{
			name: "Get sum of all bets",
			lottery: types.Lottery{
				EnterLotteryTxs: []*types.EnterLotteryTx{
					{Bet: 1}, {Bet: 2}, {Bet: 3}, {Bet: 4}, {Bet: 5}, {Bet: 6},
				},
			},
			expected: 21 + 6*types.EnterLotteryGas,
		},
		{
			name: "Lottery has only one user",
			lottery: types.Lottery{
				EnterLotteryTxs: []*types.EnterLotteryTx{{Bet: 3}},
			},
			expected: 3 + types.EnterLotteryGas,
		},
		{
			name:     "Lottery without users",
			lottery:  types.Lottery{},
			expected: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Action
			actual := c.lottery.GetSumOfAllBetsPlusFees()

			// Assert
			require.Equal(t, c.expected, actual)
		})
	}
}

func TestGetValidAddress(t *testing.T) {
	// SetUp
	expected, _ := sdk.AccAddressFromBech32(alice)

	// Action
	actual, err := types.GetAddress(alice)

	// Assert
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestGetAnotherValidAddress(t *testing.T) {
	// SetUp
	expected, _ := sdk.AccAddressFromBech32(bob)

	// Action
	actual, err := types.GetAddress(bob)

	// Assert
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestGetInvalidAddress(t *testing.T) {
	// SetUp
	_, err := sdk.AccAddressFromBech32("invalid")
	expected := sdkerrors.Wrapf(err, types.ErrInvalidUserAddress.Error(), "invalid")

	// Action
	actual, err := types.GetAddress("invalid")

	// Assert
	require.Nil(t, actual)
	require.Equal(t, expected.Error(), err.Error())
}
