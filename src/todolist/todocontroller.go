package todolist

import (
	"encoding/json"
	"firego/src/common/kv/client"
	"firego/src/common/util"
	"net/http"

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

// Todo: 目前只支持一个用户
var user_id = "agedbfqz9f5s"

func (ctl *TodoController) AddTodo(c *gin.Context) {
	type AddTodoReq struct {
		Todo string `form:"todo" json:"todo" binding:"required"`
	}
	req := &AddTodoReq{}
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error("bind json failed, err", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	id := util.GetSnowflake().String()
	data, err := json.Marshal(&TodoModel{
		Id:       id,
		Name:     req.Todo,
		Finished: false,
	})
	if err != nil {
		logrus.Error("json.marshal failed, err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	key := id
	value := string(data)

	logrus.Info(key, value)

	ctl.DB.Put(user_id, key, value)

	c.JSON(http.StatusOK, value)
}

func (ctl *TodoController) GetTodo(c *gin.Context) {
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

	data, err := json.Marshal(todo_list)
	if err != nil {
		logrus.Error("json.marshal failed, err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, string(data))
}

func (ctl *TodoController) RemoveTodo(c *gin.Context) {
	type RemoveTodoReq struct {
		Id string `form:"id" json:"id" binding:"required"`
	}
	req := &RemoveTodoReq{}
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error("bind json failed, err", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	ctl.DB.Delete(user_id, req.Id)

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
	logrus.Info()

	var data []byte
	data, err = json.Marshal(todo)
	if err != nil {
		logrus.Error("marshal failed, ", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	logrus.Info(string(data))

	ctl.DB.Put(user_id, req.Id, string(data))

	c.JSON(http.StatusOK, string(data))
}
