/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:01:34
 * @Descripttion:
 * @LastEditTime: 2020-10-16 20:19:15
 */
package util

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	twitter_epoch int64 = 1288834974657
)

//Min min
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//CheckError CheckErrorc
func CheckError(err error) {
	if err != nil {
		logrus.Info("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func CheckAndResponseError(err error, c *gin.Context) bool {
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return true
	}

	return false
}

func GetSnowflake() snowflake.ID {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logrus.Error("snowflake error ", err.Error())
		os.Exit(1)
	}

	return node.Generate()
}

// Snowflake2Unix 雪花算法ID -> unix时间戳(ms)
func Snowflake2Unix(snowflakeId string) int64 {
	id, _ := strconv.ParseInt(snowflakeId, 10, 64)

	unixtime := id>>22 + twitter_epoch

	return unixtime
}

func IsBefore1Day(snowflakeId string) bool {
	now := int64(time.Now().Unix())
	return (now - Snowflake2Unix(snowflakeId)/1000) > 86400
}
