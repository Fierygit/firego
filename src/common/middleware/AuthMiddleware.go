package middleware

/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:42:15
 * @Descripttion:
 * @LastEditTime: 2020-10-16 10:03:49
 */

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware a
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		fmt.Println("token check: " + tokenString)

		// 用户存在 将user 的信息写入上下文
		ctx.Set("user", "you")

		ctx.Next()
	}
}
