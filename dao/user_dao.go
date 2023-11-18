package dao

import (
	"errors"
	"regexp"
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

func (u *UserDAO) Update(user *model.User) error {
	updateMap := map[string]any{
		"email":         user.Email,
		"phone":         user.Phone,
		"country":       user.Country,
		"gender":        user.Gender,
		"qualification": user.Qualification,
	}
	err := u.DB.Model(&model.User{}).Where("id = ?", user.Id).Updates(updateMap).Error
	if err != nil {
		log.Error("update user invokes error", zap.Error(err))
		return errors.New("fail to update user")
	}
	return nil
}

func (u *UserDAO) GetById(id int64) (*model.User, error) {
	var user model.User
	if err := u.DB.Where("id =?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDAO) GetByUsername(username string) (*model.User, error) {
	var user model.User

	// Injected sql: xbw' and 0 in (select sleep(15) ) --

	// Apply the regex for only alphanumeric characters to prevent sql injection
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if regex.MatchString(username) {
		if err := u.DB.Where("username =?", username).First(&user).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid username")
	}
	return &user, nil
}

func (u *UserDAO) CountByEmail(id int64, email string) int64 {
	var count int64
	u.DB.Model(&model.User{}).Where("id != ? and email =?", id, email).Count(&count)
	return count
}
