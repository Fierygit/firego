/*
 * @Author: Firefly
 * @Date: 2021-03-29 13:03:59
 * @Descripttion:
 * @LastEditTime: 2021-04-23 15:48:47
 */
package todolist

import (
	mid "firego/src/common/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// Run port: ":8716"
func Run(port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	time.Now()

	todo_controller := NewTodoController()

	r.Use(mid.CORSMiddleware())
	r.Use(func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Set("user_id", "test")
		}
	}())
	r.GET("/todo", todo_controller.GetTodo)
	r.POST("/todo", todo_controller.AddTodo)
	r.POST("/todo/delete", todo_controller.RemoveTodo)
	r.POST("/todo/finish", todo_controller.FinishTodo)
	r.POST("/todo/edit", todo_controller.EditTodo)

	r.Run(port)
}
