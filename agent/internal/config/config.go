package config

import (
	"Linda/baselibs/abstractions/xlog"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lukaproject/xerr"
)

var (
	logger = xlog.NewForPackage()
)

type Config struct {
	AgentCentralEndPoint string `xdefault:"localhost:5883" xenv:"LINDA_AGENT_CENTRAL_ENDPOINT"`
	NodeId               string `xdefault:"" xenv:"LINDA_NODE_ID"`
	LocalDBDir           string `xdefault:"/tmp/linda-agent/db" xenv:"LINDA_LOCAL_DB_DIR"`
	HeartbeatPeriodMs    int    `xdefault:"50" xenv:"LINDA_HB_PERIOD_MS"`
	NodeName             string `xdefault:"test-node-name" xenv:"LINDA_NODE_NAME"`
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
	if other.NodeName != "" {
		c.NodeName = other.NodeName
	}
}

func (c *Config) getNodeId() {
	url := c.AgentAPIUrl("http") + "/agent/innercall/nodeidgen"
	resp := xerr.Must(http.Get(url))
	c.NodeId = string(xerr.Must(io.ReadAll(resp.Body)))
	logger.Infof("get node id from service, id = %s", c.NodeId)
}

func (c *Config) SetupNodeId() {
	if c.NodeId == "" {
		logger.Warn("didn't set env variable LINDA_NODE_ID yet, ask for node id from service.")
		c.getNodeId()
	}
}
