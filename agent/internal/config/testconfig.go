package config

import (
	"math/rand"
	"strconv"
)

func TestConfig() *Config {
	return &Config{
		AgentCentralUrlPrefix: "ws://localhost:5883/api/agent/heartbeat/",
		BagName:               "test-bagname",
		NodeId:                "test-bagname-" + strconv.Itoa(rand.Int()),
	}
}
