package keeper

import (
	"librarychain/x/librarychain/types"
)

var _ types.QueryServer = Keeper{}
