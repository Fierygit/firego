package middleware

/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:42:15
 * @Descripttion:
 * @LastEditTime: 2020-11-18 13:33:03
 */

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware r
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
