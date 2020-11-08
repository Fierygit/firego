/*
 * @Author: Firefly
 * @Date: 2020-09-14 19:08:16
 * @Descripttion:
 * @LastEditTime: 2020-11-07 21:24:58
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

	// 加载静态资源
	router.Static("/beibei/2020", "beibei/frontend/birthday2020")
	router.Static("/beibei/lovetree", "beibei/frontend/lovetree")

	router.Run(":2222")

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
