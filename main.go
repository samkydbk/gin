package main

import (
	"myweb/core"
	"myweb/routers"
)

func main() {
	core.Connection()
	router := routers.InitRouter()
	//静态资源
	router.Run(":8081")
}
