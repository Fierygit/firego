/*
 * @Author: Firefly
 * @Date: 2021-03-29 13:03:59
 * @Descripttion:
 * @LastEditTime: 2021-04-23 15:48:47
 */
package todolist

import (
	mid "firego/src/common/middleware"

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

	todo_controller := NewTodoController()

	r.Use(mid.CORSMiddleware())
	r.Use(mid.AuthMiddleware())
	r.GET("/todo", todo_controller.GetTodo)
	r.POST("/todo", todo_controller.AddTodo)
	r.POST("/todo/delete", todo_controller.RemoveTodo)
	r.POST("/todo/finish", todo_controller.FinishTodo)
	r.POST("/todo/edit", todo_controller.EditTodo)

	r.GET("/todo/daily/:id", todo_controller.GetDailyTodo)
	r.POST("/todo/daily", todo_controller.PutDailyTodo)

	r.Run(port)
}
