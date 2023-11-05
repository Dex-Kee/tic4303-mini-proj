package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserPageSet = wire.NewSet(wire.Struct(new(UserPage), "*"))

type UserPage struct {
}

func (u *UserPage) LoginPage(c *gin.Context) {
	title := "Login Page"
	c.HTML(http.StatusOK, "login.html", gin.H{"title": title})
}

func (u *UserPage) HomePage(c *gin.Context) {
	title := "Home Page"
	c.HTML(http.StatusOK, "home.html", gin.H{"title": title})
}

func (u *UserPage) LogoutPage(c *gin.Context) {
	title := "Logout Page"
	c.HTML(http.StatusOK, "logout.html", gin.H{"title": title})
}
