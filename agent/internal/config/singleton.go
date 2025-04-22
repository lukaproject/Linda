package config

import (
	"Linda/baselibs/abstractions/ds"
	"Linda/baselibs/abstractions/xconfig"
	"Linda/baselibs/abstractions/xos"
	"encoding/json"
	"flag"

	"github.com/lukaproject/xerr"
)

var (
	c          *Config = nil
	configFile *string = flag.String("f", "", "config file")
)

func Instance() *Config {
	return c
}

func Initial() {
	SetInstance(xconfig.NewFromOSEnv[Config]())
	if *configFile != "" {
		currentConfig := &Config{}
		xerr.Must0(json.Unmarshal(xos.ReadBytesFromFile(*configFile), currentConfig))
		c.Merge(currentConfig)
	}
	Instance().SetupNodeId()
}

func GetOSEnvs() ds.Set[string] {
	return xconfig.GetEnvs[*Config]()
}
