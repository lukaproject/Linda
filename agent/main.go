package main

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/handler"
	"Linda/agent/internal/localdb"
	"Linda/baselibs/abstractions/xconfig"
	"Linda/baselibs/abstractions/xlog"
)

func main() {
	config.SetInstance(xconfig.NewFromOSEnv[config.Config]())
	config.Instance().SetupNodeId()
	xlog.Initial()
	localdb.Initial()
	data.Initial()
	h := handler.NewHandler(nil)
	h.Run()
}
