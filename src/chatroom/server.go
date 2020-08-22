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

	// 加载静态资源
	router.Static("/static", "../static")

	// 动态路由
	v1 := router.Group("/v1")
	v1.GET("/", api.Index)
	v1.GET("/board", api.CreateBoard)
	v1.GET("/board/:id", api.EnterBoard)
	v1.GET("/message", api.GetMessage)
	v1.POST("/message", api.CreateMessage)
	v1.POST("/createroom", CreateRoom)

	router.Run(":8080")

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
