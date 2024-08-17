package config

import "fmt"

type Config struct {
	AgentCentralEndPoint string `xdefault:"localhost:5883" xenv:"LINDA_AGENT_CENTRAL_ENDPOINT"`
	BagName              string `xdefault:"testbag" xenv:"LINDA_BAG_NAME"`
	NodeId               string `xdefault:"testnodeid-1" xenv:"LINDA_NODE_ID"`
}

func (c *Config) AgentCentralUrl() string {
	return fmt.Sprintf("ws://%s/api/agent/heartbeat/%s", c.AgentCentralEndPoint, c.NodeId)
}

func (c *Config) Merge(other *Config) {
	if other.AgentCentralEndPoint != "" {
		c.AgentCentralEndPoint = other.AgentCentralEndPoint
	}
	if other.BagName != "" {
		c.BagName = other.BagName
	}
	if other.NodeId != "" {
		c.NodeId = other.NodeId
	}
}
