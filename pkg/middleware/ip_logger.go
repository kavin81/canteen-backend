package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func IPLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("[middleware] does not do anything for now")
		c.Next()
	}
}
