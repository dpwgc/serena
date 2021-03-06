package main

import (
	"fmt"
	"serena/config"
	"serena/router"
	"serena/server"
)

func main() {
	config.InitConfig()
	server.InitLog()
	server.InitRegistry()

	//初始化路由
	r := router.InitRouters()

	//获取端口号
	port := config.Get.Server.Port
	err := r.Run(fmt.Sprintf("%s%s", ":", port))
	if err != nil {
		panic(err)
	}
}
