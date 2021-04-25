package todolist

import (
	"firego/src/common/kv/client"

	"github.com/vmihailenco/msgpack"
)

/*
打卡的时候，要考虑用户重复打卡和取消打卡的情况，所以我觉得用定时器的方法比较好
每天晚上12点，扫描所有用户的todolist，
对标记了star并且finisheed的todo进行record，然后还要把todo的finished改回为false (edited)
*/

type TodoModel struct {
	Id       string
	Name     string
	Finished bool //记录是是否当天完成了
	Daily    bool
}

type TodoDailyModel struct {
	Id      string
	Records []string // etc. 2021-04-25
}

///////////////// TodoModel ////////////////////

type TodoCRUD struct {
	db client.Leveldb
}

func NewTodoCRUD() TodoCRUD {
	db := client.NewConnector().SetSize(2).Connect(client.PRE_TODO, "123456")
	return TodoCRUD{db: db}
}

func (crud *TodoCRUD) Add(user_id, todo_id, name string, finished, daily bool) (TodoModel, error) {
	todo := TodoModel{
		Id:       todo_id,
		Name:     name,
		Finished: finished,
		Daily:    daily,
	}

	data, err := msgpack.Marshal(todo)
	if err != nil {
		return TodoModel{}, err
	}

	crud.db.Put(user_id, todo_id, string(data))

	return todo, nil
}

func (crud *TodoCRUD) Delete(user_id, todo_id string) {
	crud.db.Delete(user_id, todo_id)
}

func (crud *TodoCRUD) Update(user_id, todo_id string, newTodo TodoModel) error {
	var data []byte
	data, err := msgpack.Marshal(newTodo)
	if err != nil {
		return err
	}

	crud.db.Put(user_id, todo_id, string(data))

	return nil
}

func (crud *TodoCRUD) Get(user_id, todo_id string) (TodoModel, error) {
	var todo TodoModel
	payload := crud.db.Get(user_id, todo_id)

	err := msgpack.Unmarshal([]byte(payload), &todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (crud *TodoCRUD) BatchGet(user_id string) ([]TodoModel, error) {
	todos := crud.db.BatchGet(user_id)

	todo_list := make([]TodoModel, 0)

	for _, t := range todos {
		var todo TodoModel
		err := msgpack.Unmarshal([]byte(t), &todo)
		if err != nil {
			return todo_list, err
		}
		todo_list = append(todo_list, todo)
	}

	return todo_list, nil
}

func (crud *TodoCRUD) BatchGetAll() ([]TodoModel, error) {
	todos := crud.db.BatchGetAll()

	todo_list := make([]TodoModel, 0)

	for _, t := range todos {
		var todo TodoModel
		err := msgpack.Unmarshal([]byte(t), &todo)
		if err != nil {
			return todo_list, err
		}
		todo_list = append(todo_list, todo)
	}

	return todo_list, nil
}

///////////////// TodoModel ////////////////////

///////////////// TodoDailyModel ////////////////////
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

///////////////// TodoDailyModel ////////////////////
