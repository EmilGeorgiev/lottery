package keeper

import (
	"context"
	"strconv"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const maxBet = 100

func (k msgServer) EnterLottery(goCtx context.Context, msg *types.MsgEnterLottery) (*types.MsgEnterLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Bet > maxBet {
		return nil, types.ErrExceedMaxBet
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
