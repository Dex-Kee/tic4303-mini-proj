package pojo

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	jwt.RegisteredClaims
	UserId int64  `json:"userId"`
	Role   string `json:"role"`
}
