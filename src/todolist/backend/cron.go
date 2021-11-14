package todolist

import (
	"firego/src/user"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func dateFormate(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}

func CheckDailyTodo() {
	user_crud := user.NewUserCRUD()
	todo_crud := NewTodoCRUD()
	todo_daily_crud := NewTodoDailyCRUD()

	yesterday := time.Now().AddDate(0, 0, -1)
	yesterday_format := dateFormate(yesterday)

	logrus.Info(yesterday_format)

	logrus.Info("------------------- check daily todo start ------------------")
	users_list, err := user_crud.BatchGet()
	if err != nil {
		logrus.Error(err)
		return
	}

	for _, u := range users_list {

		todos_list, err := todo_crud.BatchGet(u.Uid)
		if err != nil {
			logrus.Error(err)
			return
		}

		for _, todo := range todos_list {
			if todo.Daily && todo.Finished {
				// reset finish flag
				todo.Finished = false
				err := todo_crud.Update(u.Uid, todo.Id, todo)
				if err != nil {
					logrus.Error(err)
				}

				// add record
				todo_daily := todo_daily_crud.Get(u.Uid, todo.Id)
				new_records := append(todo_daily.Records, yesterday_format)
				todo_daily_crud.Add(u.Uid, todo.Id, new_records)
			}
		}

	}

	logrus.Info("------------------- check daily todo finished ------------------")
}
