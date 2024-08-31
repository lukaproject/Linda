package config

func TestConfig() *Config {
	return &Config{
		AgentCentralEndPoint: "localhost:5883",
		BagName:              "{bags uuid}",
		NodeId:               "test-bag-nodeid-1",
	}
}

func SetInstance(conf *Config) {
	c = conf
}
