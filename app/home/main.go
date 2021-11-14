package main

import (
	home "firego/app/home/server"
	_ "firego/comm/log" // 初始化logrus
)

func main() {
	home.Run(":6666")
}
