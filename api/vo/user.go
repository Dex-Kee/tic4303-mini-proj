package vo

type UserVO struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Country       string `json:"country"`
	Qualification string `json:"qualification"`
	Gender        string `json:"gender"`
}
