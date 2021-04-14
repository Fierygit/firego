package todolist

import (
	mid "firego/src/common/middleware"

	"github.com/gin-gonic/gin"
)

// Run port: ":8716"
func Run(port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	todo_controller := NewTodoController()

	r.Use(mid.CORSMiddleware(), mid.AuthMiddleware())
	r.GET("/todo", todo_controller.GetTodo)
	r.POST("/todo", todo_controller.AddTodo)
	r.POST("/todo/delete", todo_controller.RemoveTodo)
	r.POST("/todo/finish", todo_controller.FinishTodo)
	r.POST("/todo/edit", todo_controller.EditTodo)

	r.Run(port)
}
