package chatroom

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Index 默认首页
func Index(c *gin.Context) {
	logrus.Info("index")
	c.JSON(http.StatusOK, gin.H{
		"name": "fire go",
	})
}

// CreateBoard 创建留言板
func CreateBoard(c *gin.Context) {
	id := uuid.New()
	// TODO(创建留板言的逻辑)
	redirectURL := fmt.Sprintf("http://127.0.0.1:8080/v1/board/%s", id)
	c.Redirect(http.StatusFound, redirectURL)
}