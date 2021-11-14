package ddl

import (
	"firego/app/ddl/controller"
	mid "firego/comm/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//Run run
func Run() {
	initConfig()
	r := gin.Default()

	r.Use(mid.CORSMiddleware(), mid.RecoveryMiddleware())

	r.GET("ddl/hello", controller.Hello)

	ddlController := controller.NewDdlController()

	group := r.Group("/ddl/user")
	group.Use(mid.AuthMiddleware())
	group.POST("/add", ddlController.AddUser)
	group.DELETE(":id", ddlController.Delete)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
