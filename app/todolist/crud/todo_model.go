package todolist

import (
	"firego/server/kv/client"

	"github.com/vmihailenco/msgpack"
)

type TodoModel struct {
	Id       string
	Name     string
	Finished bool //记录是是否当天完成了
	Daily    bool
}

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
