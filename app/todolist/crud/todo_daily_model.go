package todolist

import (
	"firego/server/kv/client"

	"github.com/vmihailenco/msgpack"
)

/*
打卡的时候，要考虑用户重复打卡和取消打卡的情况，所以我觉得用定时器的方法比较好
每天晚上12点，扫描所有用户的todolist，
对标记了star并且finisheed的todo进行record，然后还要把todo的finished改回为false (edited)
*/

type TodoDailyModel struct {
	Id      string
	Records []string // etc. 2021-04-25
}

type TodoDailyCRUD struct {
	db client.Leveldb
}

func NewTodoDailyCRUD() TodoDailyCRUD {
	db := client.NewConnector().SetSize(2).Connect(client.PRE_TODO_DAILY, "123456")
	return TodoDailyCRUD{db: db}
}

func (crud *TodoDailyCRUD) Add(user_id, todo_id string, records []string) (TodoDailyModel, error) {
	todo_daily := TodoDailyModel{Id: todo_id, Records: records}
	data, err := msgpack.Marshal(todo_daily)
	if err != nil {
		return todo_daily, err
	}

	crud.db.Put(user_id, todo_id, string(data))

	return todo_daily, nil
}

func (crud *TodoDailyCRUD) Delete(user_id, todo_id string) {
	crud.db.Delete(user_id, todo_id)
}

func (crud *TodoDailyCRUD) Has(user_id, todo_id string) bool {
	return crud.db.Has(user_id, todo_id)
}

func (crud *TodoDailyCRUD) Get(user_id, todo_id string) TodoDailyModel {
	var todo_daily TodoDailyModel

	payload := crud.db.Get(user_id, todo_id)
	msgpack.Unmarshal([]byte(payload), &todo_daily) // no need to check error

	if todo_daily.Records == nil {
		todo_daily.Records = []string{}
	}

	return todo_daily
}
