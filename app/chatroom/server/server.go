package chatroom

import (
	chatroom "firego/app/chatroom/controller"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Run 启动http服务器
func Run(port string) {
	logrus.Info("start to run server!!!")
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	controller := chatroom.NewChatRoomController()

	// 加载静态资源
	router.Static("/static", "../static")
	router.StaticFile("/favicon.ico", "../static/favicon.ico")

	// 动态路由
	v1 := router.Group("/v1")
	v1.GET("/", chatroom.Index)
	v1.POST("/room", controller.CreateRoom)
	v1.POST("/enterroom", controller.EnterRoom)
	v1.POST("/message", controller.CreateMessage)
	v1.GET("/message", controller.GetMessage)

	router.Run(port)

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
