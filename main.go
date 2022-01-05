package main

import (
	"serena/config"
	"serena/server"
)

func main() {
	config.InitConfig()
	server.InitRegistry()
}
