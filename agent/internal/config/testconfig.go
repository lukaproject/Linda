package config

func TestConfig() *Config {
	return &Config{
		AgentCentralUrlFormat: "ws://localhost:5883/api/agent/heartbeat/%s",
		BagName:               "test-bagname",
	}
}
