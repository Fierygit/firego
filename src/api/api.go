package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Index 默认首页
func Index(c *gin.Context) {
	logrus.Info("index")
	c.JSON(200, gin.H{
		"name": "mingor",
	})
}

// CreateBoard 创建留言板
func CreateBoard(c *gin.Context) {
	id := uuid.New()
	logrus.Info(id)

	c.JSON(200, gin.H{
		"uuid": id,
	})
}

// Pig 写给🐖文的实例
func Pig(c *gin.Context){
	params := c.Params
	logrus.Info(params)
}