package types_test

import (
	"testing"

	"github.com/EmilGeorgiev/lottery/testutil/sample"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgEnterLottery_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgEnterLottery
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgEnterLottery{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgEnterLottery{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
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

func TestGetBetCoin(t *testing.T) {
	// SetUp
	m := types.MsgEnterLottery{
		Bet:   5,
		Denom: "token",
	}

	// Action
	actual := m.GetBetCoin()

	// Assert
	expected := sdk.NewCoin("token", sdk.NewInt(5))
	require.Equal(t, expected, actual)
}
