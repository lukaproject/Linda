package main

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/handler"
	"Linda/agent/internal/localdb"
	"Linda/baselibs/abstractions/xconfig"
)

func main() {
	config.SetInstance(xconfig.NewFromOSEnv[config.Config]())
	localdb.Initial()
	h := handler.NewHandler(nil)
	h.Run()
}
