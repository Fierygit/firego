package user

import (
	mid "firego/src/common/middleware"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	user_controller := NewUserController()

	r.Use(mid.CORSMiddleware())
	r.StaticFile("/user/login.html", "./user/static/login.html")
	r.POST("/user/login", user_controller.Login)

	r.Run(port)
}
