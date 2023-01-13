package keeper

import (
	"context"
	"fmt"
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

	lottery.RegisterNewUser(msg)
	k.Keeper.SetLottery(ctx, lottery)

	ctx.GasMeter().ConsumeGas(types.EnterLotteryGas, "Enter lottery")

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
		return fmt.Errorf(types.ErrExceedMaxBet.Error(), msg.Bet, maxBet)
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
