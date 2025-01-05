package config

import (
	"Linda/baselibs/abstractions/xconfig"
	"Linda/baselibs/abstractions/xos"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/lukaproject/xerr"
)

var c = xconfig.NewFromOSEnv[Config]()

func Instance() *Config {
	return c
}

func Initial(configfile string) {
	// read file if file exist.
	if xos.PathExists(configfile) && !xos.IsDir(configfile) {
		fromfile := Config{}
		xerr.Must0(
			json.Unmarshal(xerr.Must(io.ReadAll(
				xerr.Must(os.Open(configfile)),
			)), &fromfile))
		c.Merge(&fromfile)
	}

	if strings.ToLower(c.Env) == "test" {
		c.FileSaver = &FileSaverConfig{
			RootDir: xos.CurrentPath(),
		}
	}
}
