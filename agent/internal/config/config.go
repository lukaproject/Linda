package config

import (
	"fmt"
	"time"
)

type Config struct {
	AgentCentralEndPoint string `xdefault:"localhost:5883" xenv:"LINDA_AGENT_CENTRAL_ENDPOINT"`
	NodeId               string `xdefault:"testnodeid-1" xenv:"LINDA_NODE_ID"`
	LocalDBDir           string `xdefault:"/tmp/linda-agent/db" xenv:"LINDA_LOCAL_DB_DIR"`
	HeartbeatPeriodMs    int    `xdefault:"50" xenv:"LINDA_HB_PERIOD_MS"`
}

func (c *Config) AgentHeartBeatUrl() string {
	return fmt.Sprintf("ws://%s/api/agent/heartbeat/%s", c.AgentCentralEndPoint, c.NodeId)
}

func (c *Config) HeartbeatPeriod() time.Duration {
	return time.Duration(c.HeartbeatPeriodMs) * time.Millisecond
}

func (c *Config) AgentAPIUrl(protocol string) string {
	return protocol + "://" + c.AgentCentralEndPoint + "/api"
}

func (c *Config) Merge(other *Config) {
	if other.AgentCentralEndPoint != "" {
		c.AgentCentralEndPoint = other.AgentCentralEndPoint
	}
	if other.NodeId != "" {
		c.NodeId = other.NodeId
	}
	if other.LocalDBDir != "" {
		c.LocalDBDir = other.LocalDBDir
	}
	if other.HeartbeatPeriodMs != 0 {
		c.HeartbeatPeriodMs = other.HeartbeatPeriodMs
	}
}
