package lottery

import (
	"github.com/EmilGeorgiev/lottery/x/lottery/keeper"
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetLottery(ctx, genState.Lottery)
	k.SetSystemInfo(ctx, genState.SystemInfo)

	// Set all the finishedLottery
	for _, elem := range genState.FinishedLotteryList {
		k.SetFinishedLottery(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all lottery
	lottery, found := k.GetLottery(ctx)
	if found {
		genesis.Lottery = lottery
	}
	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.FinishedLotteryList = k.GetAllFinishedLottery(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
