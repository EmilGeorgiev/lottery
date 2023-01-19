package keeper

import (
	"fmt"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DeductFeeDecorator deducts tax by a given fee rate.
// The tax is sent to the fee_collector module.
// Call next AnteHandler if tax successfully sent.
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator.
type DeductFeeDecorator struct {
	bk types.BankKeeper
}

// NewDeductFeeDecorator initialize and return a new DeductFeeDecorator.
func NewDeductFeeDecorator(bk types.BankKeeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		bk: bk,
	}
}

// AnteHandle is a custom implementation of the AnteHandle. It deducts fees from the sender account when MsgEnterLottery
// transaction is sent. Also, it validates the fees and if the amount is not enough an error is returned.
//
// NOTE: if this handler is used you must not use the default deduct handler, because the client will be charged twice.
// By default, the ante handle deduct fees from the sender account to the fee collector. For now, we don't want that,
// because the fees can be paid as reward to the winner of the lottery.
func (dtd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if len(tx.GetMsgs()) > 0 {
		_, ok := tx.GetMsgs()[0].(*types.MsgEnterLottery)
		if !ok {
			// if the tx is different from types.MsgEnterLottery we don't want
			// to charge the sender. There is no requirements for that.
			return next(ctx, tx, simulate)
		}
	}

	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	f := feeTx.GetFee()
	if f.IsZero() {
		return sdk.Context{}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFee, fmt.Sprintf("Tx must contains exactly %d fee", types.TxFee))
	}

	if f[0].Amount.Int64() != types.TxFee {
		return sdk.Context{}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFee, fmt.Sprintf("Tx must contains exactly %d fee", types.TxFee))
	}

	// Send fees to the module
	tax := sdk.NewCoin("token", sdk.NewInt(5))
	if err = dtd.bk.SendCoinsFromAccountToModule(ctx, feeTx.FeePayer(), types.ModuleName, sdk.Coins{tax}); err != nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	events := sdk.Events{sdk.NewEvent(sdk.EventTypeTx,
		sdk.NewAttribute(sdk.AttributeKeyFee, "5"),
	)}
	ctx.EventManager().EmitEvents(events)

	return next(ctx, tx, simulate)
}
