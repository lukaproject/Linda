package config

func TestConfig() *Config {
	return &Config{
		AgentCentralEndPoint: "localhost:5883",
		NodeId:               "test-bag-nodeid-1",
		LocalDBDir:           "/tmp/linda-agent/db",
	}
}

func SetInstance(conf *Config) {
	c = conf
}
