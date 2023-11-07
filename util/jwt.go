package util

import (
	"errors"
	"tic4303-mini-proj/api/pojo"

	"github.com/golang-jwt/jwt/v4"
)

func ParseToken(token string, signingKey []byte) (*pojo.JwtCustomClaims, error) {
	if token == "" {
		return nil, errors.New("token is not found")
	}

	jwtToken, err := jwt.ParseWithClaims(token, &pojo.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(*pojo.JwtCustomClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}
