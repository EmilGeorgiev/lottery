package keeper

import (
	"fmt"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// PayReward is used to pay reward to the winner. The behavior of the method is as follows:
//   - If the winner placed the highest bet the entire pool is paid to the winner (including fees).
//   - if the winner paid the lowest bet, no reward is given, lottery total pool is carried over.
//   - All other results: winner is paid the sum of all bets (without fees) in the current lottery only
func (k *Keeper) PayReward(ctx sdk.Context, winnerIndex int64, si *types.SystemInfo, lottery types.Lottery) uint64 {
	lowestBet, highestBet := lottery.GetLowestAndHighestBet()

	// If the winner placed the highest bet the entire pool
	// (not only from current lottery) is paid to the winner (including fees)
	if (lowestBet != highestBet) && (lottery.EnterLotteryTxs[winnerIndex].Bet == highestBet) {
		currentLotteryBetsPlusFees := lottery.GetSumOfAllBetsPlusFees()
		addr, _ := types.GetAddress(lottery.EnterLotteryTxs[winnerIndex].UserAddress)
		entirePool := currentLotteryBetsPlusFees + si.LotteryPool
		coin := sdk.NewCoin(lottery.EnterLotteryTxs[winnerIndex].Denom, sdk.NewInt(int64(entirePool)))
		if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin)); err != nil {
			panic(fmt.Sprintf(types.ErrCannotPayRewards.Error(), err.Error()))
		}
		si.LotteryPool = 0
		return entirePool
	}

	// If the winner doesn't wager the highest bet the fees will not be paid as part of the
	// reward, and we can move them from the lottery module to the fee_collector module.
	fees := len(lottery.EnterLotteryTxs) * types.TxFee
	coins := sdk.NewCoin(lottery.EnterLotteryTxs[winnerIndex].Denom, sdk.NewInt(int64(fees)))
	if err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.FeeCollectorName, sdk.NewCoins(coins)); err != nil {
		panic(fmt.Sprintf(types.ErrCannotSendFeesToCollector.Error(), err.Error()))
	}

	if (lowestBet != highestBet) && (lottery.EnterLotteryTxs[winnerIndex].Bet == lowestBet) {
		// if the winner paid the lowest bet, no reward is given, lottery total pool is carried over
		si.LotteryPool += lottery.GetSumOfAllBets()
		return 0
	}

	// All other results: winner is paid the sum of all bets (without fees) in the current lottery only
	allBets := lottery.GetSumOfAllBets()
	addr, _ := types.GetAddress(lottery.EnterLotteryTxs[winnerIndex].UserAddress)
	coin := sdk.NewCoin(lottery.EnterLotteryTxs[winnerIndex].Denom, sdk.NewInt(int64(allBets)))
	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin)); err != nil {
		panic(fmt.Sprintf(types.ErrCannotPayRewards.Error(), err.Error()))
	}
	return allBets
}

// CollectBet transfer coins from user's account to the module. It is used when a user wants to enter the lottery.
// The method returns an error if the user address is invalid or sends coins from account to the module failed.
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

// RefundBet refund coins to the user's account. The method is used when the user sends more than one valid
// transactions. The lottery must keep only the last bet and the old bet must be refunded to the user. The
// panic of the user's address is invalid or the coins can't be send from the module to the user.
func (k *Keeper) RefundBet(ctx sdk.Context, tx *types.EnterLotteryTx) {
	addr, err := types.GetAddress(tx.UserAddress)
	if err != nil {
		panic(err.Error())
	}
	coin := sdk.NewCoin(tx.Denom, sdk.NewInt(int64(tx.Bet)))
	if err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin)); err != nil {
		panic(fmt.Sprintf(types.ErrCannotPayRewards.Error(), err.Error()))
	}
}
