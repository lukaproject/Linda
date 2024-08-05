package config

type Config struct {
	AgentCentralUrlPrefix string
	BagName               string
	NodeId                string
}

func (c *Config) AgentCentralUrl() string {
	return c.AgentCentralUrlPrefix + c.NodeId
}

func (c *Config) Merge(other *Config) {
	if other.AgentCentralUrlPrefix != "" {
		c.AgentCentralUrlPrefix = other.AgentCentralUrlPrefix
	}
	if other.BagName != "" {
		c.BagName = other.BagName
	}
	if other.NodeId != "" {
		c.NodeId = other.NodeId
	}
}
