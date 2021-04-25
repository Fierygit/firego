package todolist

import (
	"firego/src/common/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todo_crud       TodoCRUD
	todo_daily_crud TodoDailyCRUD
}

func NewTodoController() TodoController {
	return TodoController{
		todo_crud:       NewTodoCRUD(),
		todo_daily_crud: NewTodoDailyCRUD(),
	}
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
	todo, err := ctl.todo_crud.AddTodo(user_id, id, req.Todo, false, false)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctl *TodoController) GetTodo(c *gin.Context) {
	user_id := getUserId(c)
	todo_list, err := ctl.todo_crud.BatchGetTodo(user_id)
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

	// 顺便删除 todo daily
	ctl.todo_crud.DeleteTodo(user_id, req.Id)
	ctl.todo_daily_crud.DeleteTodoDaily(user_id, req.Id)

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
	oldTodo, err := ctl.todo_crud.GetTodo(user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Finished = req.Finished

	err = ctl.todo_crud.UpdateTodo(user_id, req.Id, newTodo)
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
	oldTodo, err := ctl.todo_crud.GetTodo(user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Name = req.Todo

	err = ctl.todo_crud.UpdateTodo(user_id, req.Id, newTodo)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, newTodo)
}

func (ctl *TodoController) GetDailyTodo(c *gin.Context) {
	type DailyTodoReq struct {
		Id string `form:"id" json:"id" binding:"required"`
	}
	req := &DailyTodoReq{}
	err := c.BindJSON(&req)
	if util.CheckAndResponseError(err, c) {
		return
	}

	user_id := getUserId(c)

	todo_daily := ctl.todo_daily_crud.GetTodoDaily(user_id, req.Id)
	todo, err := ctl.todo_crud.GetTodo(user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id":      todo.Id,
		"todo":    todo.Name,
		"records": todo_daily.Records,
	})
}

func (ctl *TodoController) PutDailyTodo(c *gin.Context) {
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
	oldTodo, err := ctl.todo_crud.GetTodo(user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Daily = req.Daily

	err = ctl.todo_crud.UpdateTodo(user_id, req.Id, newTodo)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, newTodo)
}
