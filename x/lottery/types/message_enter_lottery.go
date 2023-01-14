package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnterLottery = "enter_lottery"

var _ sdk.Msg = &MsgEnterLottery{}

func NewMsgEnterLottery(creator string, bet uint64, denom string) *MsgEnterLottery {
	return &MsgEnterLottery{
		Creator: creator,
		Bet:     bet,
		Denom:   denom,
	}
}

func (m *MsgEnterLottery) Route() string {
	return RouterKey
}

func (m *MsgEnterLottery) Type() string {
	return TypeMsgEnterLottery
}

func (m *MsgEnterLottery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgEnterLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgEnterLottery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (m *MsgEnterLottery) GetAddress() (address sdk.AccAddress, err error) {
	address, err = sdk.AccAddressFromBech32(m.Creator)
	return address, sdkerrors.Wrapf(err, ErrInvalidUserAddress.Error(), m.Creator)
}

func (m *MsgEnterLottery) GetBetCoin() (wager sdk.Coin) {
	return sdk.NewCoin(m.Denom, sdk.NewInt(int64(m.Bet)))
}
