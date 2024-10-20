package main

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/handler"
	"Linda/agent/internal/localdb"
	"Linda/baselibs/abstractions/xconfig"
)

func main() {
	config.SetInstance(xconfig.NewFromOSEnv[config.Config]())
	config.Instance().Setup()
	localdb.Initial()
	data.Initial()
	h := handler.NewHandler(nil)
	h.Run()
}
