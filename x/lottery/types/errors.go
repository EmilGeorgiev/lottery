package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrExceedMaxBet = sdkerrors.Register(ModuleName, 1100, "The bet: %d exceed the maximum allowed: %d")
)
