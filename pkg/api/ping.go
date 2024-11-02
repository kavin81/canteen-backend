package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterPingRoute(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong!",
		})
	})
}
