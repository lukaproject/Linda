package config

import "Linda/baselibs/abstractions/defaultor"

var c *Config = defaultor.New[Config]()

func Instance() *Config {
	return c
}
