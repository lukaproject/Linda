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
		PGSQL_DSN: "host=localhost user=dxyinme password=123456 dbname=linda port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		Port:      5883,
	}
}
