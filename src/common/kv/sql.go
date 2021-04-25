/*
 * @Author: Firefly
 * @Date: 2020-10-16 19:50:49
 * @Descripttion:
 * @LastEditTime: 2020-10-16 21:45:09
 */
package kv

import (
	"firego/src/common/kv/client"
	"firego/src/common/util"
	"firego/src/todolist"
	"firego/src/user"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

//TestClient test
func TestClient(what, user_id string) {
	switch what {
	case "test":
		testDB()
	case "user":
		getAllUser()
	case "todo":
		getAllTodo(user_id)
	case "opt":
		getAllOpt(user_id)
	case "json2msg":
		json2msg()
	default:
		logrus.Warn("invalid input")
	}
}

func json2msg() {
	// user_db := client.NewConnector().SetSize(2).Connect(client.PRE_USER, "123456")
	// todo_db := client.NewConnector().SetSize(2).Connect(client.PRE_TODO, "123456")

	// users_payload := user_db.BatchGet("user")
	// for i, u_payload := range users_payload {
	// 	var u user.UserModel
	// 	err := json.Unmarshal([]byte(u_payload), &u)
	// 	if err != nil {
	// 		logrus.Warn(err)
	// 		return
	// 	}

	// 	msg_payload, err := msgpack.Marshal(u)
	// 	if err != nil {
	// 		logrus.Warn(err)
	// 		return
	// 	}
	// 	logrus.Info(i, u)

	// 	user_db.Put("user", u.Name, string(msg_payload))

	// 	todos_payload := todo_db.BatchGet(u.Uid)
	// 	for i, t_payload := range todos_payload {
	// 		var t todolist.TodoModel
	// 		err := json.Unmarshal([]byte(t_payload), &t)
	// 		if err != nil {
	// 			logrus.Warn(err)
	// 			return
	// 		}

	// 		msg_payload, err := msgpack.Marshal(t)
	// 		if err != nil {
	// 			logrus.Warn(err)
	// 			return
	// 		}
	// 		logrus.Info(i, t)

	// 		todo_db.Put(u.Uid, t.Id, string(msg_payload))
	// 	}
	// }
}

func testDB() {
	// user_id := util.GetSnowflake().Base36()
	user_id := "agedbfqz9f5s"
	leveldb := client.NewConnector().SetSize(2).Connect(client.PRE_TEST, "123456")

	logrus.Info(user_id)

	key := ""
	for i := 0; i < 10; i++ {
		key = util.GetSnowflake().String()
		leveldb.Put(user_id, key, strconv.Itoa(i))
	}

	time.Sleep(time.Second * 3)

	for i, t := range leveldb.BatchGet(user_id) {
		logrus.Info(i, t)
	}
}

func getAllUser() {
	user_crud := user.NewUserCRUD()
	user_list, _ := user_crud.BatchGetUser()

	for i, u := range user_list {
		logrus.Info(i, u)
	}
}

func getAllTodo(user_id string) {
	todo_crud := todolist.NewTodoCRUD()

	todo_list, _ := todo_crud.BatchGetTodo(user_id)

	for i, todo := range todo_list {
		logrus.Info(i, todo)
	}
}

func getAllOpt(user_id string) {
	leveldb := client.NewConnector().SetSize(2).Connect(client.PRE_OPT, "123456")

	for i, todo := range leveldb.BatchGet(user_id) {
		logrus.Info(i, todo)
	}
}
