package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RegisterNewTx register a new user in the lottery. If
// the user already exist then the last bet is counted.
func (m *Lottery) RegisterNewTx(msg *MsgEnterLottery) *EnterLotteryTx {
	elt := &EnterLotteryTx{
		UserAddress: msg.Creator,
		Bet:         msg.Bet,
		Denom:       msg.Denom,
		Datetime:    FormatDateTime(time.Now()),
	}

	for i, tx := range m.EnterLotteryTxs {
		if tx.UserAddress == elt.UserAddress {
			// if the same user has new lottery transactions, then only the last
			// one counts, counter doesn’t increase on substitution.
			old := tx
			m.EnterLotteryTxs[i] = elt
			return old
		}
	}
	m.EnterLotteryTxs = append(m.EnterLotteryTxs, elt)
	return nil
}

func (m *Lottery) GetLowestAndHighestBet() (uint64, uint64) {
	lowest := m.EnterLotteryTxs[0].Bet
	highest := m.EnterLotteryTxs[0].Bet
	for _, u := range m.EnterLotteryTxs[1:] {
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
	for _, u := range m.EnterLotteryTxs {
		result += u.Bet
	}
	return
}

func (m *Lottery) GetSumOfAllBetsPlusFees() uint64 {
	fees := uint64(len(m.EnterLotteryTxs) * EnterLotteryGas)
	return m.GetSumOfAllBets() + fees
}

func GetAddress(addr string) (address sdk.AccAddress, err error) {
	address, err = sdk.AccAddressFromBech32(addr)
	return address, sdkerrors.Wrapf(err, ErrInvalidUserAddress.Error(), addr)
}
