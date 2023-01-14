package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RegisterNewUser register a new user in the lottery. If
// the user already exist then the last bet is counted.
func (m *Lottery) RegisterNewUser(msg *MsgEnterLottery) {
	u := &User{
		Address: msg.Creator,
		Bet:     msg.Bet,
	}

	for i, user := range m.Users {
		if user.Address == u.Address {
			// if the same user has new lottery transactions, then only the last
			// one counts, counter doesn’t increase on substitution.
			m.Users[i] = u
			return
		}
	}

	m.Users = append(m.Users, u)
}

func (m *Lottery) GetLowestAndHighestBet() (uint64, uint64) {
	lowest := m.Users[0].Bet
	highest := m.Users[0].Bet
	for _, u := range m.Users[1:] {
		if u.Bet < lowest {
			lowest = u.Bet
		}
		if u.Bet > highest {
			highest = u.Bet
		}
	}
	return lowest, highest
}

func (m *Lottery) GetSumOfAllBets() (result uint64) {
	for _, u := range m.Users {
		result += u.Bet
	}
	return
}

func (m *Lottery) GetSumOfAllBetsPlusFee() uint64 {
	fees := uint64(len(m.Users) * EnterLotteryGas)
	return m.GetSumOfAllBets() + fees
}

func (msg MsgEnterLottery) GetAddress() (address sdk.AccAddress, err error) {
	address, err = sdk.AccAddressFromBech32(msg.Creator)
	return address, sdkerrors.Wrapf(err, ErrInvalidUserAddress.Error(), msg.Creator)
}

func (msg *MsgEnterLottery) GetBetCoin() (wager sdk.Coin) {
	return sdk.NewCoin(msg.Denom, sdk.NewInt(int64(msg.Bet)))
}

func GetAddress(addr string) (address sdk.AccAddress, err error) {
	address, err = sdk.AccAddressFromBech32(addr)
	return address, sdkerrors.Wrapf(err, ErrInvalidUserAddress.Error(), addr)
}
