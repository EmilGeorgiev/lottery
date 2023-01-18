package keeper

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gogo/protobuf/proto"
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

	// Check whether the proposer has transaction in the lottery.
	proposerAddr := sdk.ConsAddress(ctx.BlockHeader().ProposerAddress)
	for _, tx := range lottery.EnterLotteryTxs {
		accAddr, err := types.GetAddress(tx.UserAddress)
		if err != nil {
			panic(err.Error())
		}
		acc := k.accaunt.GetAccount(ctx, accAddr)
		consAddr := sdk.GetConsAddress(acc.GetPubKey())
		if proposerAddr.Equals(consAddr) {
			// the chosen block proposer can't have any lottery transactions with itself as a sender,
			// if this is the case, then the lottery won’t fire this block, and continue on the next one
			return
		}
	}

	winnerIndex := getWinnerIndex(lottery)
	si, found := k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
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

func getWinnerIndex(l types.Lottery) int64 {
	data, _ := proto.Marshal(&l)
	b := md5.Sum(data)
	hex := fmt.Sprintf("%x", b[len(b)-2:])
	n, _ := strconv.ParseInt(hex, 16, 64)
	return n % int64(len(l.EnterLotteryTxs))
}
