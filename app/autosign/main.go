package main

import (
	autosign "firego/app/autosign/server"
	_ "firego/comm/log"
)

func main() {
	autosign.Run(":9528")
}
