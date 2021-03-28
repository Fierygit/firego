package todolist

import (
	mid "firego/src/common/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	initConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	todo_controller := NewTodoController()

	r.Use(mid.CORSMiddleware())
	r.GET("/todo", todo_controller.GetTodo)
	r.POST("/todo", todo_controller.AddTodo)
	r.POST("/todo/delete", todo_controller.RemoveTodo)
	r.POST("/todo/finish", todo_controller.FinishTodo)

	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	}
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/todolist/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
