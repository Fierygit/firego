/*
 * @Author: Firefly
 * @Date: 2020-09-14 19:08:16
 * @Descripttion:
 * @LastEditTime: 2020-11-17 09:25:56
 */
package beibei

import (
	_ "firego/src/common/log" // 初始化logrus
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
	router.NoRoute(PageNotfound())
	router.Static("/beibei/2020", "beibei/frontend/birthday2020")
	router.Static("/beibei/love", "beibei/frontend/lovetree")
	router.GET("/beibei/api/data", GetTest)
	router.GET("/beibei/api/search", SearchData)
	router.GET("/beibei/api/wordset", GetWordSetData)
	router.Run(":2222")

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
