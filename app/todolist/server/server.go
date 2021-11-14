package todolist

import (
	todolist "firego/app/todolist/controller"
	mid "firego/comm/middleware"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

// Run port: ":8716"
func Run(port string) {
	c := cron.New()
	c.AddFunc("30 0 * * *", CheckDailyTodo) // every day 00:30.am
	c.Start()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	todo_controller := todolist.NewTodoController()

	r.Use(mid.CORSMiddleware())
	r.Use(mid.AuthMiddleware())
	r.GET("/todo", todo_controller.GetTodo)
	r.POST("/todo", todo_controller.AddTodo)
	r.POST("/todo/delete", todo_controller.RemoveTodo)
	r.POST("/todo/finish", todo_controller.FinishTodo)
	r.POST("/todo/edit", todo_controller.EditTodo)

	r.GET("/todo/daily/:id", todo_controller.GetDailyTodo)
	r.POST("/todo/daily", todo_controller.ToggleDailyTodo)

	r.Run(port)
}
