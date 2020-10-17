/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:42:15
 * @Descripttion:
 * @LastEditTime: 2020-10-16 11:43:23
 */
package middleware

import (
	"firego/src/common/response"
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
