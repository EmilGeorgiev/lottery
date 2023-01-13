package keeper

import (
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetFinishedLottery set a specific finishedLottery in the store from its index
func (k Keeper) SetFinishedLottery(ctx sdk.Context, finishedLottery types.FinishedLottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinishedLotteryKeyPrefix))
	b := k.cdc.MustMarshal(&finishedLottery)
	store.Set(types.FinishedLotteryKey(
		finishedLottery.Index,
	), b)
}

// GetFinishedLottery returns a finishedLottery from its index
func (k Keeper) GetFinishedLottery(
	ctx sdk.Context,
	index string,

) (val types.FinishedLottery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinishedLotteryKeyPrefix))

	b := store.Get(types.FinishedLotteryKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFinishedLottery removes a finishedLottery from the store
func (k Keeper) RemoveFinishedLottery(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinishedLotteryKeyPrefix))
	store.Delete(types.FinishedLotteryKey(
		index,
	))
}

// GetAllFinishedLottery returns all finishedLottery
func (k Keeper) GetAllFinishedLottery(ctx sdk.Context) (list []types.FinishedLottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinishedLotteryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FinishedLottery
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
