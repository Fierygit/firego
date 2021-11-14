package user

import (
	user "firego/app/user/controller"
	mid "firego/comm/middleware"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	user_controller := user.NewUserController()

	r.Use(mid.CORSMiddleware())
	r.StaticFile("/user/login.html", "../static/login.html")
	r.POST("/user/login", user_controller.Login)

	r.Run(port)
}
