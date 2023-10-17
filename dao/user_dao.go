package dao

import (
	"errors"
	"tic4303-mini-proj/model"

	log "github.com/dzhcool/sven/zapkit"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserDAO), "*"))

type UserDAO struct {
	DB *gorm.DB
}

func (u *UserDAO) Create(user *model.User) error {
	err := u.DB.Create(user).Error
	if err != nil {
		log.Error("create user invokes error", zap.Error(err))
		return errors.New("fail to create user")
	}
	return nil
}

func (u *UserDAO) GetByUsername(username string) *model.User {
	var user model.User
	u.DB.Where("username =?", username).First(&user)
	return &user
}
