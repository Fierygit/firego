package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// EnterRoom 进入房间
func EnterRoom(ctx *gin.Context) {
	// TODO
	id := ctx.Param("id")
	logrus.Info(id)
	ctx.JSON(http.StatusOK, gin.H{"url":"fd"})
}
