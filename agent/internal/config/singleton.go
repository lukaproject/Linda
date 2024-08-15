package config

var c *Config = defaultConfig()

func Instance() *Config {
	return c
}

func Initial() {
	// read config from os env
}

func defaultConfig() *Config {
	return &Config{
		AgentCentralUrlPrefix: "ws://localhost:5883/api/agent/heartbeat/",
	}
}
