package api

import (
	"fmt"
	"net/http"
	"tic4303-mini-proj/api/dto"
	"tic4303-mini-proj/api/vo"
	"tic4303-mini-proj/service"
	"tic4303-mini-proj/util/req"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserApiSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserSvc *service.UserSvc
}

func (u *UserApi) Login(c *gin.Context) {
	var loginReq dto.LoginReq
	err := c.ShouldBind(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp("request body is not found"))
		return
	}

	fmt.Printf("form: %v\n", loginReq)

	token, err := u.UserSvc.Login(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.SuccessResp(token))
}

func (u *UserApi) Logout(c *gin.Context) {
	token := c.GetHeader("token")
	u.UserSvc.Logout(token)
	c.JSON(http.StatusOK, vo.SuccessResp(token))
}

func (u *UserApi) Profile(c *gin.Context) {
	userVO, err := u.UserSvc.Profile(webctx.GetUserId(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp(err.Error()))
		return
	}
	c.JSON(http.StatusOK, vo.SuccessResp(userVO))
}

func (u *UserApi) Create(c *gin.Context) {
	var userCreateReq dto.UserCreateReq
	err := c.ShouldBindJSON(&userCreateReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp("request body is not found"))
		return
	}

	err = u.UserSvc.Create(&userCreateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.ErrorResp(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, vo.SuccessResp(""))
}

func (u *UserApi) Update(c *gin.Context) {
	var userUpdateReq dto.UserUpdateReq
	err := c.ShouldBindJSON(&userUpdateReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.BadRequestResp("request body is not found"))
		return
	}

	userUpdateReq.Id = webctx.GetUserId(c)
	fmt.Printf("form: %v\n", userUpdateReq)

	err = u.UserSvc.Update(&userUpdateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.ErrorResp(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, vo.SuccessResp(""))
}
