package router

import (
	"tic4303-mini-proj/api"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.Struct(new(ServerRouter), "*")

type ServerRouter struct {
	UserApi *api.UserApi
}

func (s *ServerRouter) RegisterApi(app *gin.Engine) {
	userGroup := app.Group("/api/user")
	{
		userGroup.POST("/login", s.UserApi.Login)
		userGroup.PUT("/create", s.UserApi.Create)
	}
}

func (s *ServerRouter) RegisterPage(app *gin.Engine) {
	userGroup := app.Group("/page/user")
	{
		userGroup.GET("/login", s.UserApi.Login)
	}
}

func (s *ServerRouter) NoRouterHandler(c *gin.Context) {

}

func (s *ServerRouter) NoMethodHandler(c *gin.Context) {

}
