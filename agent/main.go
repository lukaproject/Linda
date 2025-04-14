package main

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/handler"
	"Linda/agent/internal/localdb"
	"Linda/baselibs/abstractions/xlog"
	"flag"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	config.Initial()
	xlog.Initial()
	localdb.Initial()
	data.Initial()
	h := handler.NewHandler(nil)
	h.Run()
}
