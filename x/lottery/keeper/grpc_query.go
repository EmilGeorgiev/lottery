package keeper

import (
	"github.com/EmilGeorgiev/lottery/x/lottery/types"
)

var _ types.QueryServer = Keeper{}
