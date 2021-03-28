/*
 * @Author: Firefly
 * @Date: 2020-10-16 19:50:49
 * @Descripttion:
 * @LastEditTime: 2020-10-16 21:45:09
 */
package client

import (
	"firego/src/common/util"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

//TestClient test
func TestClient() {

	// user_id := util.GetSnowflake().Base36()
	user_id := "agedbfqz9f5s"
	leveldb := NewConnector().SetSize(2).Connect(PRE_TEST, "123456")

	logrus.Info(user_id)

	key := ""
	for i := 0; i < 10; i++ {
		key = util.GetSnowflake().String()
		leveldb.Put(user_id, key, strconv.Itoa(i))
	}

	time.Sleep(time.Second * 3)

	logrus.Info(leveldb.BatchGet(user_id))
}
