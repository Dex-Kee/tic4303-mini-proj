package service

import (
	"tic4303-mini-proj/dao"

	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSvc), "*"))

type UserSvc struct {
	UserDAO *dao.UserDAO
}
