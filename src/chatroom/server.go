package chatroom

import (
	"firego/src/api"
	_ "firego/src/log" // 初始化logrus
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Run 启动http服务器
func Run() {
	logrus.Info("start to run server!!!")
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.GET("/", api.Index)
	router.GET("/board", api.CreateBoard)
	router.GET("/board/:id", api.EnterBoard)

	router.Run(":8080")
	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
