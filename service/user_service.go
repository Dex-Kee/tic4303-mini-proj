package service

import (
	"errors"
	"tic4303-mini-proj/api/dto"
	"tic4303-mini-proj/api/vo"
	"tic4303-mini-proj/dao"
	"tic4303-mini-proj/model"
	"tic4303-mini-proj/util"
	"time"

	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSvc), "*"))

type UserSvc struct {
	UserDAO *dao.UserDAO
}

func (u *UserSvc) Login(form *dto.LoginReq) (*vo.UserVO, error) {
	// find by username
	user := u.UserDAO.GetByUsername(form.Username)
	if user == nil {
		return nil, errors.New("username does not exist")
	}

	// check password correctness TODO
	if util.DigestSHA256(form.Password+user.PasswordSalt) != user.Password {
		return nil, errors.New("password mismatch")
	}

	return &vo.UserVO{
		Username: user.Username,
		Age:      user.Age,
		School:   user.School,
	}, nil
}

func (u *UserSvc) Create(form *dto.UserCreateReq) error {
	// check username uniqueness
	user := u.UserDAO.GetByUsername(form.Username)
	if user != nil {
		return errors.New("username already exists")
	}

	salt := util.RandomString(8)

	// create user
	u.UserDAO.Create(&model.User{
		Username:     form.Username,
		Password:     util.DigestSHA256(form.Password + salt),
		PasswordSalt: salt,
		Age:          form.Age,
		School:       form.School,
		CreateTime:   time.Now(),
	})
	
	return nil
}
