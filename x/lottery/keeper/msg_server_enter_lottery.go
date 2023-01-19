package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const maxBet = 100

func (k msgServer) EnterLottery(goCtx context.Context, msg *types.MsgEnterLottery) (*types.MsgEnterLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.validate(ctx, msg); err != nil {
		return nil, err
	}

	lottery, found := k.Keeper.GetLottery(ctx)
	if !found {
		// if we cannot find the Lottery object we call panic, because there is no way to
		// continue if it is not there. It is not like a user error, which would warrant
		// returning an error
		panic("Lottery not found")
	}

	addr, err := msg.GetAddress()
	if err != nil {
		return nil, err
	}

	oldUser := lottery.RegisterNewTx(msg)
	if err = k.bank.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(msg.GetBetCoin())); err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrCanNotPayBet.Error())
	}
	if oldUser != nil {
		// if the user already has a valid transaction in the lottery
		// we must refund the old bet and keep only the last one.
		k.Keeper.RefundBet(ctx, oldUser)
	}
	k.Keeper.SetLottery(ctx, lottery)

	//ctx.GasMeter().ConsumeGas(types.EnterLotteryGas, "Enter lottery")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EnterLotteryEventType,
			sdk.NewAttribute(types.EnterLotteryEventUser, msg.Creator),
			sdk.NewAttribute(types.EnterLotteryEventBet, strconv.FormatUint(msg.Bet, 10)),
		),
	)

	return &types.MsgEnterLotteryResponse{}, nil
}

func (k msgServer) validate(ctx sdk.Context, msg *types.MsgEnterLottery) error {
	if msg.Bet > maxBet {
		return sdkerrors.Wrap(fmt.Errorf(types.ErrExceedMaxBet.Error(), msg.Bet, maxBet), "Bet is too higher")
	}

	adr, err := msg.GetAddress()
	if err != nil {
		return err
	}
	coin := k.bank.GetBalance(ctx, adr, msg.Denom)
	if coin.Amount.Uint64() < (msg.Bet + types.EnterLotteryGas) {
		return fmt.Errorf(types.ErrNotEnoughFunds.Error(), msg.Bet+types.EnterLotteryGas, coin.Amount.Uint64())
	}

	return nil
}
