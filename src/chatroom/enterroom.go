package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type EnterRoomReq struct {
	RoomId   string `form:"roomId" binding:"required"`
	Password string `form:"password" `
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
		return
	}

	room := rooms.GetRoom(req.RoomId)
	isOpen := room.RoomInfo.CreateInfo.IsOpen
	Password := room.RoomInfo.CreateInfo.Password
	if isOpen == 0 && req.Password != Password {
		ctx.JSON(http.StatusOK, gin.H{"msg": "password error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": room})
}
