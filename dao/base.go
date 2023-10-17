package dao

import (
	"github.com/google/wire"
)

var DaoSet = wire.NewSet(UserSet)
