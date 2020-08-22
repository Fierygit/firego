package chatroom

import (
	"firego/src/api"
	_ "firego/src/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Run 启动http服务器
func Run() {
	logrus.Info("start to run the im server!!!")

	router := gin.Default()


	router.GET("/", api.Index)
	router.GET("/board", api.CreateBoard)
	router.GET("/pig/:name", api.Pig)


	router.POST("/createroom",CreateRoom)



	router.Run(":8080")

	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}
