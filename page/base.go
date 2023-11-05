package page

import (
	"github.com/google/wire"
)

var PageSet = wire.NewSet(UserPageSet)
