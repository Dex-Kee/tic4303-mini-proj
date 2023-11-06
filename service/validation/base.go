package validation

import "github.com/google/wire"

var ValidationSet = wire.NewSet(UserValidationSet)
