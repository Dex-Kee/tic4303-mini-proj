package dto

type UserCreateReq struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Country       string `json:"country"`
	Qualification string `json:"qualification"`
	Gender        string `json:"gender"`
}

type UserUpdateReq struct {
	Id            int64  `json:"id"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Country       string `json:"country"`
	Qualification string `json:"qualification"`
	Gender        string `json:"gender"`
}
