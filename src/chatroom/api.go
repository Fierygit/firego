package chatroom

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Index 默认首页
func Index(c *gin.Context) {
	logrus.Info("index")
	c.JSON(http.StatusOK, gin.H{
		"name": "fire go",
	})
}
