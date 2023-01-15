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

	if len(lottery.Users) < minUsersPerLottery {
		// we choose the winner only if the lottery
		// has 10 or more valid lottery transactions
		return
	}

	si, found := k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	// TODO check whether the validator has transaction.

	winnerIndex := getWinnerIndex(lottery.Users)
	reward := k.PayReward(ctx, winnerIndex, si, lottery)
	fl := types.FinishedLottery{
		Index:  strconv.FormatUint(si.NextId, 10),
		Winner: lottery.Users[winnerIndex].Address,
		Reward: reward,
	}
	k.SetFinishedLottery(ctx, fl)
	si.NextId++
	k.SetSystemInfo(ctx, si)
	lottery = types.Lottery{}
	k.SetLottery(ctx, lottery)

}

func getWinnerIndex(users []*types.User) int64 {
	data, _ := json.Marshal(users)
	hexDecimal := fmt.Sprintf("%x", md5.Sum(data))
	r, _ := strconv.ParseInt("0x"+hexDecimal, 0, 64)
	winnerIndex := (r & 0xFFFF) % int64(len(users))
	return winnerIndex
}
