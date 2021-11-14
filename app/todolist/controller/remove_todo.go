package todolist

import (
	"firego/comm/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *TodoController) RemoveTodo(c *gin.Context) {
	type RemoveTodoReq struct {
		Id string `form:"id" json:"id" binding:"required"`
	}
	user_id := getUserId(c)

	req := &RemoveTodoReq{}
	err := c.BindJSON(&req)
	if util.CheckAndResponseError(err, c) {
		return
	}

	// 顺便删除 todo daily
	ctl.todo_crud.Delete(user_id, req.Id)
	ctl.todo_daily_crud.Delete(user_id, req.Id)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
