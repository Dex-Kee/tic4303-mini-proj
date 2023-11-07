package router

import (
	"tic4303-mini-proj/api"
	"tic4303-mini-proj/middleware"
	"tic4303-mini-proj/page"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.Struct(new(ServerRouter), "*")

type ServerRouter struct {
	UserApi    *api.UserApi
	InfoApi    *api.InfoApi
	UserPage   *page.UserPage
	AuthFilter *middleware.AuthFilter
}

func (s *ServerRouter) RegisterApi(app *gin.Engine) {
	userGroup := app.Group("/api/user")
	{
		userGroup.POST("/login", s.UserApi.Login)
		userGroup.PUT("/create", s.UserApi.Create)
		userGroup.GET("/profile", s.AuthFilter.ValidateResource, s.UserApi.Profile)
		userGroup.PUT("/update", s.AuthFilter.ValidateResource, s.UserApi.Update)
		userGroup.PUT("/logout", s.AuthFilter.ValidateResource, s.UserApi.Logout)
	}
	infoGroup := app.Group("/api/info")
	{
		infoGroup.GET("/country", s.InfoApi.CountryMap)
	}
}

func (s *ServerRouter) RegisterPage(app *gin.Engine) {
	pageGroup := app.Group("/page")
	{
		userGroup := pageGroup.Group("/user")
		{
			userGroup.GET("/login", s.UserPage.LoginPage)
			userGroup.GET("/home", s.AuthFilter.ValidateResource, s.UserPage.HomePage)
			userGroup.GET("/logout", s.AuthFilter.ValidateResource, s.UserPage.LogoutPage)
		}
	}
}

func (s *ServerRouter) RegisterMiddleware(app *gin.Engine) {
	// Register dynamic analysis tool middleware.
	// Create statsviz server.
	srv, _ := statsviz.NewServer()
	ws := srv.Ws()
	index := srv.Index()
	app.GET("/debug/statsviz/*filepath", func(context *gin.Context) {
		if context.Param("filepath") == "/ws" {
			ws(context.Writer, context.Request)
			return
		}
		index(context.Writer, context.Request)
	})
}

func (s *ServerRouter) NoRouterHandler(c *gin.Context) {

}

func (s *ServerRouter) NoMethodHandler(c *gin.Context) {

}
