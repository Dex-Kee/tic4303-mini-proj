package service

import (
	"errors"
	"fmt"
	"tic4303-mini-proj/api/dto"
	"tic4303-mini-proj/api/pojo"
	"tic4303-mini-proj/api/vo"
	"tic4303-mini-proj/constant"
	"tic4303-mini-proj/dao"
	"tic4303-mini-proj/model"
	"tic4303-mini-proj/util"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSvc), "*"))

type UserSvc struct {
	UserDAO       *dao.UserDAO
	JwtSigningKey []byte
	DigestKey     string
}

func (u *UserSvc) Login(form *dto.LoginReq) (string, error) {
	// find by username
	user, err := u.UserDAO.GetByUsername(form.Username)
	if err != nil {
		return "", errors.New("username does not exist")
	}

	if util.DigestSHA256(form.Password+user.PasswordSalt) != user.Password {
		return "", errors.New("password mismatch")
	}

	// create jwt claim
	var claims = pojo.JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    constant.TokenIssuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constant.TokenValidityDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.Id,
		Role:   user.Role,
	}

	// create token
	token, err := u.createToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserSvc) Logout(token string) {
	// save to redis, mark the token as the revoked token
	fmt.Println(token)
}

func (u *UserSvc) Profile(id int64) (*vo.UserVO, error) {
	user, _ := u.UserDAO.GetById(id)
	if user == nil {
		return nil, errors.New("user does not exist")
	}
	userVO := vo.UserVO{
		Username:      user.Username,
		Email:         user.Email,
		Phone:         user.Phone,
		Country:       user.Country,
		Qualification: user.Qualification,
		Gender:        user.Gender,
	}
	return &userVO, nil
}

func (u *UserSvc) createToken(claims pojo.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.JwtSigningKey)
}

func (u *UserSvc) Create(form *dto.UserCreateReq) error {
	// check username uniqueness
	user, _ := u.UserDAO.GetByUsername(form.Username)
	if user != nil {
		return errors.New("username already exists")
	}

	salt := util.RandomString(8)

	// create user
	u.UserDAO.Create(&model.User{
		Username:      form.Username,
		Email:         form.Email,
		Phone:         form.Phone,
		Country:       form.Country,
		Qualification: form.Qualification,
		Gender:        form.Gender,
		Password:      util.DigestSHA256(form.Password + salt),
		PasswordSalt:  salt,
		Role:          constant.UserRoleStudent,
		CreateTime:    time.Now(),
	})

	return nil
}

func (u *UserSvc) Update(form *dto.UserUpdateReq) error {
	user, _ := u.UserDAO.GetById(form.Id)
	if user == nil {
		return errors.New("user does not exist")
	}

	// checking of uniqueness of field values (todo)

	user.Email = form.Email
	user.Phone = form.Phone
	user.Gender = form.Gender
	user.Country = form.Country
	user.Qualification = form.Qualification
	err := u.UserDAO.Update(user)
	if err != nil {
		return errors.New("update failed")
	}

	return nil
}

func (u *UserSvc) FindAccessibleResourceByRole(role string) map[string]bool {
	// query accessible resource by role
	// hard code for now
	resourceMap := make(map[string]bool)
	resourceMap["/api/user/profile"] = true
	resourceMap["/api/user/create"] = true
	resourceMap["/api/user/update"] = true
	resourceMap["/api/user/logout"] = true
	resourceMap["/api/user/delete"] = false
	return resourceMap
}
