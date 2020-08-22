package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func Run(){
	logrus.Info("start to run the im server!!!")
	
	r := gin.Default()
	r.POST("/createroom", func(ctx *gin.Context) {
		logrus.Info("get a con")
		logrus.Info(ctx.PostForm("k"))

		logrus.Info(ctx.Request.PostForm)
		for k, v := range ctx.Request.PostForm {
			logrus.Info("k:%v\n", k)
			logrus.Info("v:%v\n", v)
		}

		ctx.String(200, "success")
	})

	r.Run(":8080")
	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}