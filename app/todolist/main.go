package main

import (
	todolist "firego/app/todolist/server"
	_ "firego/comm/log"
)

func main() {
	todolist.Run(":8716")
}
