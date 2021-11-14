/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:42:15
 * @Descripttion:
 * @LastEditTime: 2020-11-08 13:45:36
 */
package beibei

import (
	"firego/src/common/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

//PageNotfound re
func PageNotfound() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.Error(c, fmt.Sprint("你不对劲"), nil)
		return
	}
}
