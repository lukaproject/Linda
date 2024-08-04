package config

import "fmt"

type Config struct {
	AgentCentralUrlFormat string
	BagName               string
}

func (c *Config) AgentCentralUrl() string {
	return fmt.Sprintf(c.AgentCentralUrlFormat, c.BagName)
}

func (c *Config) Merge(other *Config) {
	if other.AgentCentralUrlFormat != "" {
		c.AgentCentralUrlFormat = other.AgentCentralUrlFormat
	}
	if other.BagName != "" {
		c.BagName = other.BagName
	}
}
