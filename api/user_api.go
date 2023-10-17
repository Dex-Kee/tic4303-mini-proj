package api

import (
	"net/http"
	"tic4303-mini-proj/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSvc *service.UserSvc
}

func (u UserApi) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
