/*
 * @Author: Firefly
 * @Date: 2020-10-15 21:48:15
 * @Descripttion:
 * @LastEditTime: 2020-10-15 23:09:25
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Response data: 数据 msg: 提示
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

//Success success
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 0, data, msg)
}

//Error error
func Error(ctx *gin.Context, msg string, data gin.H) {
	Response(ctx, http.StatusOK, -1, data, msg)
}
