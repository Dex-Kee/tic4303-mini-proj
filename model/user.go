package model

import (
	"time"
)

type User struct {
	Id            int64     `gorm:"column:id"`
	Username      string    `gorm:"column:username"`
	Email         string    `gorm:"column:email"`
	Phone         string    `gorm:"column:phone"`
	Country       string    `gorm:"column:country"`
	Qualification string    `gorm:"column:qualification"`
	Gender        string    `gorm:"column:gender"`
	Password      string    `gorm:"column:password"`
	PasswordSalt  string    `gorm:"column:password_salt"`
	Role          string    `gorm:"column:role"`
	CreateTime    time.Time `gorm:"column:create_time"`
}

func (User) TableName() string {
	return "t_user"
}
