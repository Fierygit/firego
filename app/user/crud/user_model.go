package user

import (
	"firego/server/kv/client"

	"github.com/vmihailenco/msgpack"
)

const (
	kv_user_key = "user" // 因为所有用户共用一张表
)

type UserModel struct {
	Uid  string
	Name string
}

type UserCRUD struct {
	db client.Leveldb
}

func NewUserCRUD() UserCRUD {
	db := client.NewConnector().SetSize(2).Connect(client.PRE_USER, "123456")
	return UserCRUD{db: db}
}

func (user_crud *UserCRUD) Add(user_id, name string) (UserModel, error) {
	user := UserModel{
		Uid:  user_id,
		Name: name,
	}
	data, err := msgpack.Marshal(user)
	if err != nil {
		return user, err
	}
	key := name
	value := string(data)
	user_crud.db.Put(kv_user_key, key, value)

	return user, nil
}

func (user_crud *UserCRUD) Get(username string) (UserModel, error) {
	user := UserModel{}
	payload := user_crud.db.Get(kv_user_key, username)

	err := msgpack.Unmarshal([]byte(payload), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (user_crud *UserCRUD) BatchGet() ([]UserModel, error) {
	users := user_crud.db.BatchGet(kv_user_key)

	user_list := make([]UserModel, 0)

	for _, t := range users {
		var user UserModel
		err := msgpack.Unmarshal([]byte(t), &user)
		if err != nil {
			return user_list, err
		}
		user_list = append(user_list, user)
	}

	return user_list, nil
}

func (user_crud *UserCRUD) Has(username string) bool {
	return user_crud.db.Has(kv_user_key, username)
}
