package dto

import "time"

/*
 * @Author: Firefly
 * @Date: 2020-10-16 23:33:56
 * @Descripttion:
 * @LastEditTime: 2020-10-16 23:39:22
 */

//Ddl 不需要名字
type Ddl struct {
	Name       string
	CreateTime time.Time
	DeadTime   time.Time
	Title      string
	Detail     string
}
