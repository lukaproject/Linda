package main

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/handler"
	"flag"
)

var (
	configfile = flag.String("f", "etc/agent.json", "agent config file")
)

func main() {
	flag.Parse()
	config.Initial(*configfile)
	h := handler.NewHandler(config.Instance())
	h.Run()
}
