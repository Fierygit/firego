/*
 * @Author: Firefly
 * @Date: 2020-11-13 15:10:11
 * @Descripttion:
 * @LastEditTime: 2020-11-18 13:22:40
 */

/*
 * @Author: Firefly
 * @Date: 2020-09-14 19:08:16
 * @Descripttion:
 * @LastEditTime: 2020-11-08 14:59:53
 */

package home

import (
	_ "firego/src/common/log" // 初始化logrus
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Run 启动http服务器
func Run() {
	logrus.Info("start to run server!!!")
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	fmt.Println(os.Getwd())
	router.NoRoute(PageNotfound())
	router.Static("/", "home/frontend/dist/")
	router.Run(":6666")
	logrus.Info("should not run here now") // 暂时不用协成， 之后改成异步的
}

var index []byte

func readIndex() {
	b, err1 := ioutil.ReadFile("home/frontend/dist/index.html")
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	index = b
}

//PageNotfound re
func PageNotfound() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(index) == 0 || c.Request.URL.Path == "update" {
			readIndex()
		}
		c.Writer.Write(index)
		c.Writer.Flush()
		return
	}
}
