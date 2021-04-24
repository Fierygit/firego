/*
 * @Author: Firefly
 * @Date: 2021-04-23 16:34:56
 * @Descripttion:
 * @LastEditTime: 2021-04-24 16:15:24
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var f interface{}

	b := []byte(`[{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}]`)

	json.Unmarshal(b, &f)

	data, ok := f.([]interface{}) // 这里不能使用f.([]map[string]interface{})，这样是无法判断成功的

	if ok {

		fmt.Printf("%+v\n", data)
		fmt.Println()
	}
}
