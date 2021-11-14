package middleware

import (
	"firego/comm/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

//RecoveryMiddleware re
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Error(c, fmt.Sprint(err), nil)
				return
			}
		}()
		c.Next()
	}
}
