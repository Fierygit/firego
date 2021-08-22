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
	todo, err := ctl.todo_crud.Add(user_id, id, req.Todo, false, false)
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (ctl *TodoController) GetTodo(c *gin.Context) {
	query_type := c.DefaultQuery("type", "unfinished")
	user_id := getUserId(c)
	todo_list, err := ctl.todo_crud.BatchGet(user_id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	filtered_todo_list := []TodoModel{}
	daily_todos := []TodoModel{}
	switch query_type {
	case "all":
		for _, t := range todo_list {
			if t.Daily {
				daily_todos = append(daily_todos, t)
				continue
			}
			filtered_todo_list = append(filtered_todo_list, t)
		}
		ReverseTodoList(filtered_todo_list)
	case "finished":
		for _, t := range todo_list {
			if t.Finished && t.Daily {
				daily_todos = append(daily_todos, t)
				continue
			}
			if t.Finished && util.IsBeforeNDay(t.Id, 7) {
				filtered_todo_list = append(filtered_todo_list, t)
			}
		}
		ReverseTodoList(filtered_todo_list)
	case "unfinished":
		fallthrough
	default:
		for _, t := range todo_list {
			if t.Daily {
				daily_todos = append(daily_todos, t)
				continue
			}
			if !t.Finished || t.Finished && !util.IsBeforeNDay(t.Id, 7) {
				filtered_todo_list = append(filtered_todo_list, t)
			}
		}
	}

	filtered_todo_list = append(daily_todos, filtered_todo_list...)

	c.JSON(http.StatusOK, filtered_todo_list)
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
	ctl.todo_crud.Delete(user_id, req.Id)
	ctl.todo_daily_crud.Delete(user_id, req.Id)

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
	oldTodo, err := ctl.todo_crud.Get(user_id, req.Id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	newTodo := oldTodo
	newTodo.Finished = req.Finished

	err = ctl.todo_crud.Update(user_id, req.Id, newTodo)
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
