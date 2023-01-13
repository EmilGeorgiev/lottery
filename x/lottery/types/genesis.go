package types

import (
	"fmt"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	fmt.Println(".")
	return &GenesisState{
		Lottery: Lottery{},
		SystemInfo: SystemInfo{
			NextId:      1,
			LotteryPool: 0,
		},
		FinishedLotteryList: []FinishedLottery{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in finishedLottery
	finishedLotteryIndexMap := make(map[string]struct{})

	for _, elem := range gs.FinishedLotteryList {
		index := string(FinishedLotteryKey(elem.Index))
		if _, ok := finishedLotteryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for finishedLottery")
		}
		finishedLotteryIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
