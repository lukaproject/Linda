package config

import "Linda/baselibs/abstractions/xconfig"

var c *Config = xconfig.NewFromOSEnv[Config]()

func Instance() *Config {
	return c
}
