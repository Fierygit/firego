/*
 * @Author: Firefly
 * @Date: 2021-03-31 12:27:45
 * @Descripttion:
 * @LastEditTime: 2021-03-31 14:25:18
 */

package main

import (
	"firego/src/opt"
	"fmt"
	"os"
)

var usage = `
	firego backend by mingor & firefly
	Usage:
		opt [cmd] [arguments]

	Example:
		opt add user_info key
		opt get user_info
		`

func main() {

	fmt.Println(os.Args[0])
	for index, i := range os.Args {
		fmt.Println(index, i, os.Args[index])
	}

	arg_len := len(os.Args)
	if arg_len <= 1 {
		fmt.Println(usage)
		return
	}
	switch os.Args[1] {
	case "add":
		if arg_len < 4 {
			fmt.Println(usage)
		}
		if success, msg := opt.AddOpt(os.Args[2], os.Args[3]); success {
			fmt.Println("success")
		} else {
			fmt.Println(msg)
		}
	case "get":
		if arg_len < 3 {
			fmt.Println(usage)
		}
		fmt.Println("opt coe : ", opt.GetOpt(os.Args[2]))
	default:
		fmt.Println(usage)

	}

}
