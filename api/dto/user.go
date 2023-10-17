package dto

type UserCreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	School   string `json:"school"`
}
