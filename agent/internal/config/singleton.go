package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/lukaproject/xerr"
)

var c *Config = defaultConfig()

func Instance() *Config {
	return c
}

func Initial(configfile string) {
	// read file if file exist.
	if _, err := os.Stat(configfile); !os.IsNotExist(err) {
		fromfile := Config{}
		json.Unmarshal(xerr.Must(io.ReadAll(
			xerr.Must(os.Open(configfile)),
		)), &fromfile)
		c.Merge(&fromfile)
	}
}

func defaultConfig() *Config {
	return &Config{
		AgentCentralUrlPrefix: "ws://localhost:5883/api/agent/heartbeat/",
	}
}
