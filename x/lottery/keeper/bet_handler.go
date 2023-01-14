package keeper

import (
	"fmt"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k *Keeper) PayReward(ctx sdk.Context, winnerIndex int64, si types.SystemInfo, lottery types.Lottery) uint64 {
	lowestBet, highestBet := lottery.GetLowestAndHighestBet()
	if lottery.Users[winnerIndex].Bet == lowestBet {
		// if the winner paid the lowest bet, no reward is given, lottery total pool is carried over
		si.LotteryPool += lottery.GetSumOfAllBets()
		k.SetSystemInfo(ctx, si)
		return 0
	}

	// If the winner placed the highest bet the entire pool
	// (not only from current lottery) is paid to the winner (including fees)
	if lottery.Users[winnerIndex].Bet == highestBet {
		currentLotteryBetsPlusFees := lottery.GetSumOfAllBetsPlusFees()
		addr, _ := types.GetAddress(lottery.Users[winnerIndex].Address)
		entirePool := currentLotteryBetsPlusFees + si.LotteryPool
		coin := sdk.NewCoin(lottery.Users[winnerIndex].Denom, sdk.NewInt(int64(entirePool)))
		if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin)); err != nil {
			panic(fmt.Sprintf(types.ErrCannotPayRewards.Error(), err.Error()))
		}
		return entirePool
	}

	// All other results: winner is paid the sum of all bets (without fees) in the current lottery only
	allBets := lottery.GetSumOfAllBets()
	addr, _ := types.GetAddress(lottery.Users[winnerIndex].Address)
	coin := sdk.NewCoin(lottery.Users[winnerIndex].Denom, sdk.NewInt(int64(allBets)))
	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin)); err != nil {
		panic(fmt.Sprintf(types.ErrCannotPayRewards.Error(), err.Error()))
	}
	return allBets
}

func (k *Keeper) CollectBet(ctx sdk.Context, el types.MsgEnterLottery) error {
	addr, err := el.GetAddress()
	if err != nil {
		return err
	}
	if err = k.bank.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(el.GetBetCoin())); err != nil {
		return sdkerrors.Wrapf(err, types.ErrCanNotPayBet.Error())
	}
	return nil
}

func (k *Keeper) RefundBet(ctx sdk.Context, u *types.User) {
	addr, _ := types.GetAddress(u.Address)
	coin := sdk.NewCoin(u.Denom, sdk.NewInt(int64(u.Bet)))
	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin)); err != nil {
		panic(fmt.Sprintf(types.ErrCannotPayRewards.Error(), err.Error()))
	}
}
