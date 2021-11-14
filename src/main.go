/*
 * @Author: Firefly
 * @Date: 2020-10-16 10:28:36
 * @Descripttion:
 * @LastEditTime: 2021-03-31 12:19:12
 */
package main

import (
	"firego/src/autosign"
	"firego/src/beibei"
	"firego/src/common/kv"
	"firego/src/common/kv/server"
	_ "firego/src/common/log"
	"firego/src/home"

	"fmt"
	"os"
)

func main() {
	// for idx, args := range os.Args {
	// 	logrus.Info("init "+strconv.Itoa(idx)+" : ", args)
	// }

	switch os.Args[1] {
	case "home", "-h":
		home.Run()
	case "leveldb", "-l":
		server.Run()
	case "testdb":
		kv.TestClient(os.Args[2], os.Args[3])
	case "beibei":
		beibei.Run()
	case "autosign":
		autosign.Run(":9528")

	default:
		usage := `
	firego backend by mingor & firefly
	Usage:
		go run main.go [arguments]

	The commands are:
	-d ddl
	-l leveldb
		`

		fmt.Println(usage)
	}
}
