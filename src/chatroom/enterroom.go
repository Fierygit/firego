package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type EnterRoomReq struct {
	RoomId   string `form:"roomId" binding:"required"`
	Password   string `form:"password" `
	OtherMsg string `form:"other" `
}

// EnterRoom 进入房间
func EnterRoom(ctx *gin.Context) {
	// TODO
	req := &EnterRoomReq{}
	err := ctx.BindJSON(&req)
	logrus.Info(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
	}


	ctx.JSON(http.StatusOK, gin.H{"info":"test"})
}
