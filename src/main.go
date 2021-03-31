/*
 * @Author: Firefly
 * @Date: 2020-10-16 10:28:36
 * @Descripttion:
 * @LastEditTime: 2021-03-31 12:19:12
 */
package main

import (
	"firego/src/beibei"
	"firego/src/common/kv/client"
	"firego/src/common/kv/server"
	_ "firego/src/common/log"
	"firego/src/home"
	"firego/src/proxy"
	"firego/src/todolist"
	"firego/src/user"
	"fmt"
	"os"
)

func main() {
	// for idx, args := range os.Args {
	// 	logrus.Info("init "+strconv.Itoa(idx)+" : ", args)
	// }

	switch os.Args[1] {
	case "proxy", "-p":
		proxy.Run()
	case "home", "-h":
		home.Run()
	case "leveldb", "-l":
		server.Run()
	case "testdb":
		client.TestClient()
	case "beibei":
		beibei.Run()
	case "todolist":
		todolist.Run(":8716")
	case "user", "-u":
		user.Run(":9527")
	case "opt":

	default:
		usage := `
	firego backend by mingor & firefly
	Usage:
		go run main.go [arguments]

	The commands are:
	-p proxy
	-d ddl
	-l leveldb
	-u user
		`

		fmt.Println(usage)
	}
}
