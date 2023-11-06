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

func (u *UserValidationSvc) UpdateUserChecker(form *dto.UserUpdateReq) error {
	isEmail := util.IsEmail(form.Email)
	if !isEmail {
		return errors.New("email format is invalid")
	}

	isPhone := util.IsPhoneNumber(form.Phone)
	if !isPhone {
		return errors.New("phone number is invalid, it should be 8 digits and start with 6, 7, 8 or 9")
	}

	country := form.Country
	codes := util.ListCountryCode()
	_, isValidCountry := lo.Find(codes, func(v string) bool {
		return v == country
	})
	if !isValidCountry {
		return errors.New("country is invalid, is should be in ISO 3166-1 alpha-2 format")
	}

	gender := form.Gender
	if gender != "M" && gender != "F" {
		return errors.New("gender is invalid")
	}

	qualification := form.Qualification
	if qualification != "high-school" && qualification != "diploma" && qualification != "bachelor" && qualification != "master" && qualification != "phd" {
		return errors.New("qualification is invalid")
	}

	// email should not be duplicated with other user
	emailCount := u.UserDAO.CountByEmail(form.Id, form.Email)
	if emailCount > 0 {
		return errors.New("email is duplicated with other user")
	}

	return nil
}
