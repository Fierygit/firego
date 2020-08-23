package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// EnterBoard 进入留言板
func EnterBoard(c *gin.Context) {
	// TODO
	id := c.Param("id")
	logrus.Info(id)
}
