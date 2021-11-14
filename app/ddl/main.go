package main

import (
	ddl "firego/app/ddl/server"
	_ "firego/comm/log"
)

func main() {
	ddl.Run()
}
