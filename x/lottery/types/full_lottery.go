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

func (msg MsgEnterLottery) GetAddress() (address sdk.AccAddress, err error) {
	address, err = sdk.AccAddressFromBech32(msg.Creator)
	return address, sdkerrors.Wrapf(err, ErrInvalidUserAddress.Error(), msg.Creator)
}
