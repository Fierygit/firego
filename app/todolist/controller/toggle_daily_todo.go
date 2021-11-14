package todolist

import (
	"firego/comm/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *TodoController) ToggleDailyTodo(c *gin.Context) {
	type DailyTodoReq struct {
		Id    string `form:"id" json:"id" binding:"required"`
		Daily bool   `form:"daily" json:"daily" binding:"required"`
	}
	req := &DailyTodoReq{}
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
	newTodo.Daily = req.Daily

	err = ctl.todo_crud.Update(user_id, req.Id, newTodo)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, newTodo)
}
