package main

import (
	_ "firego/src/log"
	"firego/src/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("fire go!!!")

	router := gin.Default()

	router.GET("/", api.Index)
	router.GET("/board", api.CreateBoard)
	router.GET("/pig/:name", api.Pig)

	router.Run(":8080")
}
