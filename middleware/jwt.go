package middleware

import (
	"errors"
	"net/http"
	"tic4303-mini-proj/api/pojo"
	"tic4303-mini-proj/api/vo"
	"tic4303-mini-proj/constant"
	"tic4303-mini-proj/service"

	log "github.com/dzhcool/sven/zapkit"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var AuthFilterSet = wire.NewSet(wire.Struct(new(AuthFilter), "*"))

type AuthFilter struct {
	JwtSigningKey []byte
	DigestKey     string
	UserSvc       *service.UserSvc
}

func (a *AuthFilter) ValidateResource(c *gin.Context) {
	resourcePath := c.FullPath()

	// check token
	token, _ := c.Cookie("token")

	// token is empty, redirect to login page
	if token == "" {
		c.Redirect(http.StatusFound, "/page/user/login")
		return
	}

	claims, err := a.parserToken(token)
	if err != nil {
		log.Error("error when parse token: ", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, vo.UnauthorizedResp("token is invalid"))
		return
	}

	// if role is admin, no resource check is required
	role := claims.Role
	if role == constant.UserRoleAdmin {
		c.Next()
		return
	}

	// for other roles, check accessible resource
	resourceMap := a.UserSvc.FindAccessibleResourceByRole(role)
	if ok, _ := resourceMap[resourcePath]; !ok {
		log.Infof("request uri: [%s] is denial by role of user [%s]", resourcePath, role)
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			vo.ErrorResp(constant.RespCodeInvalidResourceAccess, constant.RespMsgInvalidResourceAccess))
		return
	}

	// check if the token has been revoked

	// write to header
	c.Set(constant.AppUserIdHeader, claims.UserId)
	c.Set(constant.AppUserRoleHeader, claims.Role)
	c.Next()
}

func (a *AuthFilter) parserToken(tokenString string) (*pojo.JwtCustomClaims, error) {
	if tokenString == "" {
		return nil, errors.New("token is not found")
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &pojo.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.JwtSigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(*pojo.JwtCustomClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}
