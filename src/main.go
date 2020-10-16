/*
 * @Author: Firefly
 * @Date: 2020-10-16 10:28:36
 * @Descripttion:
 * @LastEditTime: 2020-10-16 12:04:19
 */
package main

import (
	"firego/src/common/log"
	"firego/src/ddl"
	"firego/src/proxy"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {

	log.Init()
	for idx, args := range os.Args {
		logrus.Info("init "+strconv.Itoa(idx)+" : ", args)
	}

	switch os.Args[1] {
	case "proxy":
		proxy.Run()
	case "-p":
		proxy.Run()
		break
	case "ddl":
		ddl.Run()
	case "-d":
		ddl.Run()
		break
	default:
		usage := fmt.Sprintf("\nfirego backend by mingor & firefly\n\n")
		usage += fmt.Sprintf("Usage:\n\n\tgo run main.go [arguments]\n\n")
		usage += fmt.Sprintf("The commands are:\n\n")
		usage += fmt.Sprintf("\t-p\tproxy\n\t-d\tddl")

		usage += fmt.Sprintf("\n\n")
		logrus.Info(usage)
	}
}
