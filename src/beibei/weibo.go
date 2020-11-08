/*
 * @Author: Firefly
 * @Date: 2020-11-08 14:51:57
 * @Descripttion:
 * @LastEditTime: 2020-11-08 15:22:19
 */
package beibei

import (
	"encoding/json"
	"firego/src/common/response"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gin-gonic/gin"
)

//GetTest v
func GetTest(ctx *gin.Context) {
	b, err1 := ioutil.ReadFile("beibei/weibodata/dealed.json")
	if err1 != nil {
		fmt.Println("read fail", err1)
	}
	var m interface{}

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}

	// for k, v := range m {
	// 	fmt.Println(k, ":", v)
	// 	break
	// }
	fmt.Print(reflect.TypeOf(m))
	test := m.(map[string]interface{})
	response.Success(ctx, test, "success")

}
