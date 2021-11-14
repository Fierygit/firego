package main

import (
	chatroom "firego/app/chatroom/server"
	_ "firego/comm/log" // 初始化logrus
)

func main() {
	chatroom.Run(":9090")
}
