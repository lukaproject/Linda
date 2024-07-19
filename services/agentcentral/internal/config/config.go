package config

type Config struct {
	PGSQL_DSN string
	Port      int
	Redis     *RedisConfig
}

type RedisConfig struct {
	Addrs    []string
	Db       int
	Password string
}

func (c *Config) Merge(other *Config) {
	if other.PGSQL_DSN != "" {
		c.PGSQL_DSN = other.PGSQL_DSN
	}
	if other.Port != 0 {
		c.Port = other.Port
	}
	if other.Redis != nil {
		c.Redis = other.Redis
	}
}
