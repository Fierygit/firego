/*
 * @Author: Firefly
 * @Date: 2020-09-14 19:08:16
 * @Descripttion:
 * @LastEditTime: 2020-10-15 22:34:27
 */
package chatroom

import (
	_ "firego/src/common/log" // 初始化logrus

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var rooms *ChatRoom

// Run 启动http服务器
func Run() {
	logrus.Info("start to run server!!!")
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	rooms = InitChatRoom()

	router := gin.Default()

	// 加载静态资源
	router.Static("/static", "../static")
	router.StaticFile("/favicon.ico", "../static/favicon.ico")

	// 动态路由
	v1 := router.Group("/v1")
	v1.GET("/", Index)
	v1.POST("/room", CreateRoom)
	v1.POST("/enterroom", EnterRoom)
	v1.POST("/message", CreateMessage)
	v1.GET("/message", GetMessage)

	router.Run(":80")

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
