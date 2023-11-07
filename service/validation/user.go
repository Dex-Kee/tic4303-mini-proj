package validation

import (
	"errors"
	"tic4303-mini-proj/api/dto"
	"tic4303-mini-proj/dao"
	"tic4303-mini-proj/util"

	"github.com/google/wire"
	"github.com/samber/lo"
)

var UserValidationSet = wire.NewSet(wire.Struct(new(UserValidationSvc), "*"))

type UserValidationSvc struct {
	UserDAO *dao.UserDAO
}

func (u *UserValidationSvc) CreateUserChecker(form *dto.UserCreateReq) error {
	isStrongPwd := util.IsStrongPassword(form.Password)
	if !isStrongPwd {
		return errors.New("password is not strong enough, At least one uppercase letter (A-Z). At least one lowercase letter (a-z). At least one digit (0-9).\nAt least one special character (not a letter or a digit). Minimum length of 8 characters.")
	}

	err := u.userInfoChecker(0, form.Email, form.Phone, form.Country, form.Gender, form.Qualification)
	if err != nil {
		return err
	}

	// check username uniqueness
	user, _ := u.UserDAO.GetByUsername(form.Username)
	if user != nil {
		return errors.New("username already exists")
	}

	return nil
}

func (u *UserValidationSvc) UpdateUserChecker(form *dto.UserUpdateReq) error {
	return u.userInfoChecker(form.Id, form.Email, form.Phone, form.Country, form.Gender, form.Qualification)
}

func (u *UserValidationSvc) userInfoChecker(id int64, email, phone, country, gender, qualification string) error {
	isEmail := util.IsEmail(email)
	if !isEmail {
		return errors.New("email format is invalid")
	}

	isPhone := util.IsPhoneNumber(phone)
	if !isPhone {
		return errors.New("phone number is invalid, it should be 8 digits and start with 6, 7, 8 or 9")
	}

	codes := util.ListCountryCode()
	_, isValidCountry := lo.Find(codes, func(v string) bool {
		return v == country
	})
	if !isValidCountry {
		return errors.New("country is invalid, is should be in ISO 3166-1 alpha-2 format")
	}

	if gender != "M" && gender != "F" {
		return errors.New("gender is invalid")
	}

	if qualification != "high-school" && qualification != "diploma" && qualification != "bachelor" && qualification != "master" && qualification != "phd" {
		return errors.New("qualification is invalid")
	}

	// email should not be duplicated with other user
	emailCount := u.UserDAO.CountByEmail(id, email)
	if emailCount > 0 {
		return errors.New("email is duplicated with other user")
	}

	return nil
}
