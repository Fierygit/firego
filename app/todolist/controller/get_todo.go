package todolist

import (
	crud "firego/app/todolist/crud"
	"firego/comm/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *TodoController) GetTodo(c *gin.Context) {
	query_type := c.DefaultQuery("type", "unfinished")
	user_id := getUserId(c)
	todo_list, err := ctl.todo_crud.BatchGet(user_id)
	if util.CheckAndResponseError(err, c) {
		return
	}

	filtered_todo_list := []crud.TodoModel{}
	daily_todos := []crud.TodoModel{}
	switch query_type {
	case "all":
		for _, t := range todo_list {
			if t.Daily {
				daily_todos = append(daily_todos, t)
				continue
			}
			filtered_todo_list = append(filtered_todo_list, t)
		}
		crud.ReverseTodoList(filtered_todo_list)
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
		crud.ReverseTodoList(filtered_todo_list)
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
