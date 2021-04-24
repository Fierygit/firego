package todolist

import "time"

/*
打卡的时候，要考虑用户重复打卡和取消打卡的情况，所以我觉得用定时器的方法比较好
每天晚上12点，扫描所有用户的todolist，
对标记了star并且finisheed的todo进行record，然后还要把todo的finished改回为false (edited)
*/

type TodoModel struct {
	Id       string
	Content  string
	Finished bool //记录是是否当天完成了
	Daily    bool
}

type TodoModelRecords struct {
	Id      string
	Records []time.Time
}
