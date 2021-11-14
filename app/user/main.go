package main

import (
	user "firego/app/user/server"
)

func main() {
	user.Run(":9527")
}
