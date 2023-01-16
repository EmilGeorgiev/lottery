package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrExceedMaxBet              = sdkerrors.Register(ModuleName, 1100, "the bet: %d exceed the maximum allowed: %d")
	ErrInvalidUserAddress        = sdkerrors.Register(ModuleName, 1101, "new user address is invalid: %s")
	ErrNotEnoughFunds            = sdkerrors.Register(ModuleName, 1102, "the sender has not enough funds to cover lottery fee + minimum bet. Needed: %d, available: %d")
	ErrCanNotPayBet              = sdkerrors.Register(ModuleName, 1103, "can't pay the bet")
	ErrCannotPayRewards          = sdkerrors.Register(ModuleName, 1104, "cannot pay reward to winner: %s")
	ErrCannotSendFeesToCollector = sdkerrors.Register(ModuleName, 1105, "cannot send fees the fee collector")
)
