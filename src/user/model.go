package user

import (
	"encoding/json"
	"firego/src/common/kv/client"
)

const (
	kv_user_key = "user" // 因为所有用户共用一张表
)

type UserModel struct {
	Uid  string
	Name string
}

///////////////// UserModel ////////////////////
func AddUser(db client.Leveldb, user_id, name string) (UserModel, error) {
	user := UserModel{
		Uid:  user_id,
		Name: name,
	}
	data, err := json.Marshal(user)
	if err != nil {
		return user, err
	}
	key := name
	value := string(data)
	db.Put(kv_user_key, key, value)

	return user, nil
}

func HasUser(db client.Leveldb, username string) bool {
	return db.Has(kv_user_key, username)
}

///////////////// UserModel ////////////////////
