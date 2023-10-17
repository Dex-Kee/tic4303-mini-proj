package api

import (
	"net/http"
	"tic4303-mini-proj/api/dto"
	"tic4303-mini-proj/api/vo"
	"tic4303-mini-proj/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSvc *service.UserSvc
}

func (u UserApi) Login(c *gin.Context) {
	var loginReq dto.LoginReq
	err := c.ShouldBind(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp("request body is not found"))
		return
	}

	userVO, err := u.UserSvc.Login(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, vo.SuccessResp(userVO))
}

func (u UserApi) Create(c *gin.Context) {
	var userCreateReq dto.UserCreateReq
	err := c.ShouldBindJSON(&userCreateReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp("request body is not found"))
		return
	}

	err = u.UserSvc.Create(&userCreateReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, vo.SuccessResp(""))
}
