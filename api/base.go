package api

import (
	"github.com/google/wire"
)

var ApiSet = wire.NewSet(UserApiSet, InfoApiSet)
