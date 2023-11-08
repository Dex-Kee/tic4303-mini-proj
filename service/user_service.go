package service

import (
	"context"
	"errors"
	"tic4303-mini-proj/api/dto"
	"tic4303-mini-proj/api/pojo"
	"tic4303-mini-proj/api/vo"
	"tic4303-mini-proj/constant"
	"tic4303-mini-proj/constant/exception"
	"tic4303-mini-proj/dao"
	"tic4303-mini-proj/model"
	"tic4303-mini-proj/service/validation"
	"tic4303-mini-proj/util"
	"time"

	log "github.com/dzhcool/sven/zapkit"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSvc), "*"))

type UserSvc struct {
	JwtSigningKey     []byte
	DigestKey         string
	UserDAO           *dao.UserDAO
	UserValidationSvc *validation.UserValidationSvc
	RedisClient       *redis.Client
}

func (u *UserSvc) Login(form *dto.LoginReq) (string, error) {
	user, err := u.UserValidationSvc.UserLoginChecker(form)
	if err != nil {
		return "", err
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
	log.Info("user invokes Logout", zap.String("token", token))
	if token == "" {
		return
	}

	// logout should not blocking the main process
	go func() {
		claim, err := util.ParseToken(token, u.JwtSigningKey)
		if err != nil {
			log.Error("Logout invokes error", zap.Error(err))
			return
		}
		// get the token expiration time
		seconds := claim.ExpiresAt.Sub(time.Now()).Seconds()
		if seconds <= 0 {
			log.Info("token has expired, no need to revoke")
			return
		}
		// save to redis, mark the token as the revoked token
		u.RedisClient.Set(context.Background(), constant.RedisRevokedTokenKey+":"+token, "1", time.Duration(seconds)*time.Second)
	}()
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

func (u *UserSvc) createToken(claims pojo.JwtCustomClaims) (string, *exception.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(u.JwtSigningKey)
	if err != nil {
		log.Error("failed to create token", zap.Error(err))
		return "", exception.NewError(500, "failed to create token")
	}
	return signedString, nil
}

func (u *UserSvc) Create(form *dto.UserCreateReq) error {
	err := u.UserValidationSvc.CreateUserChecker(form)
	if err != nil {
		return err
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

	err := u.UserValidationSvc.UpdateUserChecker(form)
	if err != nil {
		return err
	}

	// checking of uniqueness of field values (todo)
	user.Email = form.Email
	user.Phone = form.Phone
	user.Gender = form.Gender
	user.Country = form.Country
	user.Qualification = form.Qualification
	err = u.UserDAO.Update(user)
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

	resourceMap["/page/user/home"] = true
	resourceMap["/page/user/logout"] = true
	return resourceMap
}
