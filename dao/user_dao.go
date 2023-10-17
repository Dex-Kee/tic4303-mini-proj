package dao

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserDAO), "*"))

type UserDAO struct {
	DB *gorm.DB
}
