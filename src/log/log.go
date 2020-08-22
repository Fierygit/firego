package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
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

func init() {
	logrus.SetFormatter(new(MyFormatter))
	logrus.SetReportCaller(true)
}
