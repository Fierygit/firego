package main

import (
	"firego/src/chatroom"
	"firego/src/websocket"
)

func main() {
	go websocket.Run()
	chatroom.Run()
}
