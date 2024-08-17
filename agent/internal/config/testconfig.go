package config

import (
	"math/rand"
	"strconv"
)

func TestConfig() *Config {
	return &Config{
		AgentCentralEndPoint: "localhost:5883",
		BagName:              "test-bagname",
		NodeId:               "test-bagname-" + strconv.Itoa(rand.Int()),
	}
}
