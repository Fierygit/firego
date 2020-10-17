/*
 * @Author: Firefly
 * @Date: 2020-09-14 19:08:16
 * @Descripttion:
 * @LastEditTime: 2020-10-16 11:00:24
 */
package log

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// init log formting
type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	//fmt.Println(entry.Caller.File)
	fullPath := strings.Split(entry.Caller.File, "/")
	path := fullPath[len(fullPath)-1] + ":" + strconv.Itoa(entry.Caller.Line)
	msg := fmt.Sprintf("%s %8s %15s\t %s\n",
		timestamp,
		"["+strings.ToUpper(entry.Level.String())+"]",
		path,
		entry.Message)
	return []byte(msg), nil
}

func Init() {
	logrus.SetFormatter(new(MyFormatter))
	logrus.SetReportCaller(true)
}
