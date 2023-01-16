package keeper

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const minUsersPerLottery = 10

func (k *Keeper) ChooseWinner(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lottery, found := k.GetLottery(ctx)
	if !found {
		panic("Lottery not found")
	}

	if len(lottery.EnterLotteryTxs) < minUsersPerLottery {
		// we choose the winner only if the lottery
		// has 10 or more valid lottery transactions
		return
	}

	si, found := k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	// TODO check whether the validator has transaction.

	winnerIndex := getWinnerIndex(lottery.EnterLotteryTxs)
	reward := k.PayReward(ctx, winnerIndex, &si, lottery)
	fl := types.FinishedLottery{
		Index:           strconv.FormatUint(si.NextId, 10),
		Winner:          lottery.EnterLotteryTxs[winnerIndex].UserAddress,
		Reward:          reward,
		EnterLotteryTxs: lottery.EnterLotteryTxs,
		WinnerIndex:     uint64(winnerIndex),
	}
	k.SetFinishedLottery(ctx, fl)
	si.NextId++
	k.SetSystemInfo(ctx, si)
	lottery = types.Lottery{}
	k.SetLottery(ctx, lottery)

}

func getWinnerIndex(txs []*types.EnterLotteryTx) int64 {
	data, _ := json.Marshal(txs)
	b := md5.Sum(data)
	hex := fmt.Sprintf("%x", b[len(b)-2:])
	n, _ := strconv.ParseInt(hex, 16, 64)
	return n % int64(len(txs))
}
