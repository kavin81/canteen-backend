package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	RegisterPingRoute(router)
	RegisterHelloRoute(router)
}
