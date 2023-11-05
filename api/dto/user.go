package dto

type UserCreateReq struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `gorm:"column:school"`
	Phone         string `gorm:"column:phone"`
	Country       string `gorm:"column:country"`
	Qualification string `gorm:"column:qualification"`
	Gender        string `gorm:"column:gender"`
}

type UserUpdateReq struct {
	Id            int64  `json:"id"`
	Email         string `gorm:"column:school"`
	Phone         string `gorm:"column:phone"`
	Country       string `gorm:"column:country"`
	Qualification string `gorm:"column:qualification"`
	Gender        string `gorm:"column:gender"`
}
