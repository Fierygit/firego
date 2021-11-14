package todolist

import (
	"firego/comm/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *TodoController) AddTodo(c *gin.Context) {
	type AddTodoReq struct {
		Todo string `form:"todo" json:"todo" binding:"required"`
	}
	req := &AddTodoReq{}
	err := c.BindJSON(&req)
	if util.CheckAndResponseError(err, c) {
		return
	}

	id := util.GetSnowflake().String()
	user_id := getUserId(c)
	todo, err := ctl.todo_crud.Add(user_id, id, req.Todo, false, false)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, todo)
}
