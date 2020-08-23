package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// EnterRoom 进入留言板
func EnterRoom(c *gin.Context) {
	// TODO
	id := c.Param("id")
	logrus.Info(id)
}
