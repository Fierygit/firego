package todolist

import (
	"encoding/json"
	"firego/src/common/kv/client"
	"firego/src/common/util"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	logrus.Info(user_id)

	return user_id
}

func (ctl *TodoController) AddTodo(c *gin.Context) {
	type AddTodoReq struct {
		Content string `form:"content" json:"content" binding:"required"`
		Daily   bool   `form:"daily" json:"daily" binding:"required"`
	}
	user_id := getUserId(c)

	req := &AddTodoReq{}
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error("bind json failed, err", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	id := util.GetSnowflake().String()
	todo := TodoModel{
		Id:       id,
		Content:  req.Content,
		Finished: false,
		Daily:    req.Daily,
	}
	data, err := json.Marshal(todo)
	if err != nil {
		logrus.Error("json.marshal failed, err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	key := id
	value := string(data)

	logrus.Info(key, value)

	ctl.DB.Put(user_id, key, value)

	c.JSON(http.StatusOK, todo)
}

func (ctl *TodoController) GetTodo(c *gin.Context) {
	user_id := getUserId(c)
	todos := ctl.DB.BatchGet(user_id)

	todo_list := make([]TodoModel, 0)

	for _, t := range todos {
		var todo TodoModel
		err := json.Unmarshal([]byte(t), &todo)
		if err != nil {
			logrus.Error("Unmarshal failed, ", err)
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		todo_list = append(todo_list, todo)
	}

	// data, err := json.Marshal(todo_list)
	// if err != nil {
	// 	logrus.Error("json.marshal failed, err:", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, todo_list)
}

func (ctl *TodoController) RemoveTodo(c *gin.Context) {
	type RemoveTodoReq struct {
		Id string `form:"id" json:"id" binding:"required"`
	}
	user_id := getUserId(c)

	req := &RemoveTodoReq{}
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error("bind json failed, err", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	ctl.DB.Delete(user_id, req.Id)

	// 删除时间  有 批处理  user-id , 新建一张表
	times := ctl.DB.Get(getTimesKey(user_id), req.Id)

	times_slice := strings.Split(times, ";")

	for _, val := range times_slice {
		ctl.DB.Delete(getTimesKey(user_id), val)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func (ctl *TodoController) FinishTodo(c *gin.Context) {
	type FinishTodoReq struct {
		Id       string `form:"id" json:"id" binding:"required"`
		Finished bool   `form:"finished" json:"finished" binding:"required"`
	}
	user_id := getUserId(c)

	req := &FinishTodoReq{}
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error("bind json failed, err", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var todo TodoModel
	payload := ctl.DB.Get(user_id, req.Id)

	err = json.Unmarshal([]byte(payload), &todo)
	if err != nil {
		logrus.Error("Unmarshal failed, ", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	todo.Finished = req.Finished

	var data []byte
	data, err = json.Marshal(todo)
	if err != nil {
		logrus.Error("marshal failed, ", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	logrus.Info(string(data))

	ctl.DB.Put(user_id, req.Id, string(data))

	c.JSON(http.StatusOK, todo)
}

func (ctl *TodoController) EditTodo(c *gin.Context) {
	type EditTodoReq struct {
		Id   string `form:"id" json:"id" binding:"required"`
		Todo string `form:"todo" json:"todo" binding:"required"`
	}
	user_id := getUserId(c)

	req := &EditTodoReq{}
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error("bind json failed, err", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var todo TodoModel
	payload := ctl.DB.Get(user_id, req.Id)

	err = json.Unmarshal([]byte(payload), &todo)
	if err != nil {
		logrus.Error("Unmarshal failed, ", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	todo.Content = req.Todo

	var data []byte
	data, err = json.Marshal(todo)
	if err != nil {
		logrus.Error("marshal failed, ", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	logrus.Info(string(data))

	ctl.DB.Put(user_id, req.Id, string(data))

	c.JSON(http.StatusOK, todo)
}

func getTodayKey(user_id string) string {
	return "time-" + user_id + "-" + getTimeStr()
}

func getTimesKey(user_id string) string {
	return "time-" + user_id
}

func getTimeStr() string {
	return fmt.Sprint(time.Now().Year(), "-", time.Now().Month(), "-", time.Now().Day())
}
