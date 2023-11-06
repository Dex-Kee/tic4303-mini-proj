package api

import (
	"net/http"
	"tic4303-mini-proj/api/vo"
	"tic4303-mini-proj/util"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var InfoApiSet = wire.NewSet(wire.Struct(new(InfoApi), "*"))

type InfoApi struct {
}

func (u *InfoApi) CountryMap(c *gin.Context) {
	countryMap := util.CountryMap()
	c.JSON(http.StatusOK, vo.SuccessResp(countryMap))
}
