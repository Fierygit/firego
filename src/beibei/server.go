/*
 * @Author: Firefly
 * @Date: 2020-09-14 19:08:16
 * @Descripttion:
 * @LastEditTime: 2020-11-18 13:26:11
 */
package beibei

import (
	_ "firego/src/common/log" // 初始化logrus
	mid "firego/src/common/middleware"
	"firego/src/common/response"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Run 启动http服务器
func Run() {
	logrus.Info("start to run server!!!")
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	fmt.Println(os.Getwd())
	router.Use(mid.CORSMiddleware())
	router.NoRoute(PageNotfound())
	router.Static("/beibei/api/2020", "beibei/frontend/birthday2020")
	router.GET("/beibei/2020", func(c *gin.Context) {
		response.Error(c, "活动已过期 -_-", nil)

	})
	router.Static("/beibei/love", "beibei/frontend/lovetree")
	router.GET("/beibei/api/data", GetTest)
	router.GET("/beibei/api/search", SearchData)
	router.GET("/beibei/api/wordset", GetWordSetData)
	router.Run(":2222")

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
