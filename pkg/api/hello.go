package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterHelloRoute(router *gin.RouterGroup) {
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "world!",
		})
	})
}