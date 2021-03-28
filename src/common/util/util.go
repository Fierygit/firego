/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:01:34
 * @Descripttion:
 * @LastEditTime: 2020-10-16 20:19:15
 */
package util

import (
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
)

//Min min
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//CheckError c
func CheckError(err error) {
	if err != nil {
		logrus.Info("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func GetSnowflake() snowflake.ID {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logrus.Error("snowflake error ", err.Error())
		os.Exit(1)
	}

	return node.Generate()
}
