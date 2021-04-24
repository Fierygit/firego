package todolist

import (
	"firego/src/common/kv/client"
	"time"

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

type TodoRecordModel struct {
	Id      string
	Records []time.Time
}

///////////////// TodoModel ////////////////////

func Addtodo(db client.Leveldb, user_id, todo_id, name string, finished, daily bool) (TodoModel, error) {
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

	db.Put(user_id, todo_id, string(data))

	return todo, nil
}

func DeleteTodo(db client.Leveldb, user_id, todo_id string) {
	db.Delete(user_id, todo_id)

	//TODO 删除时间  有 批处理  user-id , 新建一张表
}

func UpdateTodo(db client.Leveldb, user_id, todo_id string, newTodo TodoModel) error {
	var data []byte
	data, err := msgpack.Marshal(newTodo)
	if err != nil {
		return err
	}

	db.Put(user_id, todo_id, string(data))

	return nil
}

func GetTodo(db client.Leveldb, user_id, todo_id string) (TodoModel, error) {
	var todo TodoModel
	payload := db.Get(user_id, todo_id)

	err := msgpack.Unmarshal([]byte(payload), &todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func BatchGetTodo(db client.Leveldb, user_id string) ([]TodoModel, error) {
	todos := db.BatchGet(user_id)

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
