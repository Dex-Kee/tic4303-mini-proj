package webctx

import (
	"tic4303-mini-proj/constant"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) int64 {
	return c.GetInt64(constant.AppUserIdHeader)
}
