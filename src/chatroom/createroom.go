package chatroom

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LoginForm struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}


func CreateRoom(ctx *gin.Context){

		form := &LoginForm{}
		// message := c.BindJSON("message")
		// nick := c.PostForm("nick")
		if ctx.BindJSON(&form) == nil {
			logrus.Info(form.User, form.Password)
			if form.User == "user" && form.Password == "password" {
				ctx.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				ctx.JSON(401, gin.H{"status": "unauthorized"})
			}
		}

}