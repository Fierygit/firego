package cock

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HelloWorld() {
	logrus.Info("helloWorld")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":"mingor",
		})
	})
	r.Run(":8080")
}

