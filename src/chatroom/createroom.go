package chatroom

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

// CreateRoomReq 请求创建聊天室参数
type CreateRoomReq struct {
	IsOpen     int    `form:"isOpen"   binding:"required"`
	RoomName   string `form:"roomName" binding:"required"`
	RoomInfo   string `form:"roomInfo" binding:"required"`
	RoomNotice string `form:"roomInfo" binding:"required"`
	Password   string `form:"password" `
}

// CreateRoom 创建留言板
func CreateRoom(ctx *gin.Context) {
	req := &CreateRoomReq{}
	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
	}

	logrus.Info(req)
	if err2 := checkArg(req); err2 != nil {
		ctx.JSON(200, gin.H{"msg": err2.Error()})
		return
	}
	uuid := createRoom(req)
	redirectURL := fmt.Sprintf("http://127.0.0.1:8080/v1/room/%s", uuid)
	ctx.Redirect(http.StatusFound, redirectURL)
}

func createRoom(req *CreateRoomReq) string {
	id := uuid.New()
	// TODO used as session
	// save to db
	logrus.Info("create a new room %s", id.String())

	return id.String()
}

func checkArg(req *CreateRoomReq) error {
	if req.IsOpen != 0 && req.IsOpen != 1 {
		return errors.New("isOpen error")
	}
	if len(req.RoomName) <= 0 {
		return errors.New("roomName is error")
	}
	if req.IsOpen == 0 && (len(req.Password) <= 0 || len(req.Password) >= 30) {
		return errors.New("password is error")
	}
	return nil
}
