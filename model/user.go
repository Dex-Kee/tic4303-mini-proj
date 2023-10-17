package model

import (
	"time"
)

type User struct {
	Id           int64     `gorm:"column:id"`
	Username     string    `gorm:"column:username"`
	Password     string    `gorm:"column:password"`
	PasswordSalt string    `gorm:"column:password_salt"`
	CreateTime   time.Time `gorm:"column:create_time"`
}

func (User) TableName() string {
	return "t_user"
}
