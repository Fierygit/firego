package todolist

import (
	"firego/src/common/kv/client"
	"firego/src/common/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	DB client.Leveldb
}

func NewTodoController() TodoController {
	db := client.NewConnector().SetSize(2).Connect(client.PRE_TODO, "123456")
	return TodoController{DB: db}
}

func getUserId(c *gin.Context) string {
	user_id := c.GetString("user_id")
	return user_id
}

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
	todo, err := Addtodo(ctl.DB, user_id, id, req.Todo, false, false)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctl *TodoController) GetTodo(c *gin.Context) {
	user_id := getUserId(c)
	todo_list, err := BatchGetTodo(ctl.DB, user_id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, todo_list)
}

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

	DeleteTodo(ctl.DB, user_id, req.Id)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func (ctl *TodoController) FinishTodo(c *gin.Context) {
	type FinishTodoReq struct {
		Id       string `form:"id" json:"id" binding:"required"`
		Finished bool   `form:"finished" json:"finished" binding:"required"`
	}
	req := &FinishTodoReq{}
	err := c.BindJSON(&req)
	if util.CheckAndResponseError(err, c) {
		return
	}

	user_id := getUserId(c)
	oldTodo, err := GetTodo(ctl.DB, user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Finished = req.Finished

	err = UpdateTodo(ctl.DB, user_id, req.Id, newTodo)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, newTodo)
}

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
	oldTodo, err := GetTodo(ctl.DB, user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Name = req.Todo

	err = UpdateTodo(ctl.DB, user_id, req.Id, newTodo)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, newTodo)
}

func (ctl *TodoController) DailyTodo(c *gin.Context) {
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
	oldTodo, err := GetTodo(ctl.DB, user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Daily = req.Daily

	err = UpdateTodo(ctl.DB, user_id, req.Id, newTodo)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, newTodo)
}
