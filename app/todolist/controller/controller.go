package todolist

import (
	crud "firego/app/todolist/crud"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todo_crud       crud.TodoCRUD
	todo_daily_crud crud.TodoDailyCRUD
}

func NewTodoController() TodoController {
	return TodoController{
		todo_crud:       crud.NewTodoCRUD(),
		todo_daily_crud: crud.NewTodoDailyCRUD(),
	}
}

func getUserId(c *gin.Context) string {
	user_id := c.GetString("user_id")
	return user_id
}
