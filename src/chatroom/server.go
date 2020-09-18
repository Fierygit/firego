package chatroom

import (
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
	router.StaticFile("/favicon.ico", "../static/favicon.ico")

	// 动态路由
	v1 := router.Group("/v1")
	v1.GET("/", Index)
	v1.POST("/room", CreateRoom)
	v1.GET("/room/:id", EnterRoom)
	v1.POST("/message", CreateMessage)
	v1.GET("/message", GetMessage)

	router.Run(":666")

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
