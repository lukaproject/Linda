package config

import (
	"Linda/baselibs/abstractions/defaultor"
	"Linda/baselibs/abstractions/xos"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/lukaproject/xerr"
)

var c *Config = defaultor.New[Config]()

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

	if strings.ToLower(c.Env) == "test" {
		c.FileSaver = &FileSaverConfig{
			RootDir: xos.CurrentPath(),
		}
	}
}
