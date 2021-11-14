package todolist

import (
	"firego/comm/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *TodoController) EditTodo(c *gin.Context) {
	type EditTodoReq struct {
		Id   string `form:"id" json:"id" binding:"required"`
		Todo string `form:"todo" json:"todo" binding:"required"`
	}
	req := &EditTodoReq{}
	err := c.BindJSON(&req)
	if util.CheckAndResponseError(err, c) {
		return
	}

	user_id := getUserId(c)
	oldTodo, err := ctl.todo_crud.Get(user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Name = req.Todo

	err = ctl.todo_crud.Update(user_id, req.Id, newTodo)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, newTodo)
}

func (ctl *TodoController) GetDailyTodo(c *gin.Context) {
	todo_id := c.Param("id")

	user_id := getUserId(c)

	todo_daily := ctl.todo_daily_crud.Get(user_id, todo_id)
	todo, err := ctl.todo_crud.Get(user_id, todo_id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id":       todo.Id,
		"todo":     todo.Name,
		"finished": todo.Finished,
		"records":  todo_daily.Records,
	})
}
