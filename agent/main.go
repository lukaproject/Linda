package main

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/data"
	"Linda/agent/internal/handler"
	"Linda/agent/internal/localdb"
	"Linda/baselibs/abstractions/xlog"
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "Environment variables:\n")
	envs := config.GetOSEnvs()
	for k := range envs {
		fmt.Fprintf(flag.CommandLine.Output(), "\t%s\n", k)
	}
}

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
