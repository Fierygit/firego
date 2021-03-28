/*
 * @Author: Firefly
 * @Date: 2020-10-16 19:50:49
 * @Descripttion:
 * @LastEditTime: 2020-10-16 21:45:09
 */
package client

import (
	"time"

	"github.com/sirupsen/logrus"
)

//TestClient test
func TestClient() {

	leveldb := NewConnector().SetSize(2).Connect("firefly", "123456")
	for i := 0; i < 10; i++ {
		leveldb.Put("test", "sdf")
	}
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		logrus.Info(leveldb.Get("test"))
	}
}
