package main

import (
	"Linda/agent/internal/config"
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
