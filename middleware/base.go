package middleware

import "github.com/google/wire"

var MiddlewareSet = wire.NewSet(InitMysql, AuthFilterSet)
