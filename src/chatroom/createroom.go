package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateRoomArg struct {
	IsOpen     int    `form:"isOpen"   binding:"required"`
	RoomName   string `form:"roomName" binding:"required"`
	RoomInfo   string `form:"roomInfo" binding:"required"`
	RoomNotice string `form:"roomInfo" binding:"required"`
	Password   string `form:"password" `
}

func CreateRoom(ctx *gin.Context) {

	arg := &CreateRoomArg{}
	// message := c.BindJSON("message")
	// nick := c.PostForm("nick")
	err := ctx.BindJSON(&arg)
	if err == nil {
		logrus.Info(arg)
		if isTrue, info := checkArg(arg); !isTrue {
			ctx.JSON(200, gin.H{"code": 2, "msg": info})
			return
		}
		uuid_ := createRoom(arg)
		ctx.JSON(200, gin.H{"code": 1, "msg": "/chatroom?uuid=" + uuid_})
	} else {
		ctx.JSON(200, gin.H{"code": 1, "msg": "required some arg but not find!!!"})
	}
}

func createRoom(arg *CreateRoomArg) string {

	id := uuid.New()
	// used as session
	// save to db
	logrus.Info("create a new room %s", id.String())

	return id.String()

}

func checkArg(arg *CreateRoomArg) (bool, string) {
	if arg.IsOpen != 0 && arg.IsOpen != 1 {
		return false, "isOpen error"
	}
	if len(arg.RoomName) <= 0 {
		return false, "roomName is error"
	}
	if arg.IsOpen == 0 && (len(arg.Password) <= 0 || len(arg.Password) >= 30) {
		return false, "password is error"
	}
	return true, ""
}
