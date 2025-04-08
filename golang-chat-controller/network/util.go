package network

import (
	"github.com/gin-gonic/gin"
	"golang-chat-controller/types"
)

func response(c *gin.Context, s int, res interface{}, data ...string) {
	c.JSON(s, types.NewRes(s, res, data...))
}
