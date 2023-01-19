package keeper_test

import (
	"testing"

	"github.com/EmilGeorgiev/lottery/x/lottery/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery/testutil"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/rosetta"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const alice = "cosmos1slwzaxnrxde9dwl73ny7c7kmca0jmr0ppna0px"

func setupAnteHandeTest(t *testing.T) (keeper.DeductFeeDecorator, *testutil.MockBankKeeper, sdk.Context, client.TxBuilder) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	d := keeper.NewDeductFeeDecorator(bankMock)
	ctx := sdk.NewContext(nil, tmproto.Header{}, false, nil)
	cdc, _ := rosetta.MakeCodec()
	txConfig := tx.NewTxConfig(cdc, tx.DefaultSignModes)
	txBuilder := txConfig.NewTxBuilder()

	return d, bankMock, ctx, txBuilder
}

func TestAnteHandle_DeductFees(t *testing.T) {
	// SetUp
	d, bankMock, ctx, txBuilder := setupAnteHandeTest(t)
	el := &types.MsgEnterLottery{
		Creator: alice,
		Bet:     6,
		Denom:   "token",
	}
	_ = txBuilder.SetMsgs(el)
	coin := sdk.NewCoin("token", sdk.NewInt(int64(types.TxFee)))
	txBuilder.SetFeeAmount(sdk.NewCoins(coin))
	sender, _ := types.GetAddress(alice)
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(coin)).
		Return(nil)

	// Acction
	_, err := d.AnteHandle(ctx, txBuilder.GetTx(), false, nextAnte)

	// Assert
	require.NoError(t, err)
}

func TestAnteHandle_TheFeesAreLessThenMinimum(t *testing.T) {
	// SetUp
	d, _, ctx, txBuilder := setupAnteHandeTest(t)
	el := &types.MsgEnterLottery{
		Creator: alice,
		Bet:     6,
		Denom:   "token",
	}
	_ = txBuilder.SetMsgs(el)
	coin := sdk.NewCoin("token", sdk.NewInt(2))
	txBuilder.SetFeeAmount(sdk.NewCoins(coin))

	// Acction
	_, err := d.AnteHandle(ctx, txBuilder.GetTx(), false, nextAnte)

	// Assert
	require.Equal(t, "Tx must contains exactly 5 fee: insufficient fee", err.Error())
}

func nextAnte(ctx sdk.Context, t sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
	return ctx, nil
}
