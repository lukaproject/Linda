package config

type Config struct {
	PGSQL_DSN string `xdefault:"host=localhost user=dxyinme password=123456 dbname=linda port=5432 sslmode=disable TimeZone=Asia/Shanghai"`
	Port      int    `xdefault:"5883"`
	Redis     *RedisConfig
}

type RedisConfig struct {
	Addrs    []string `xdefault:"localhost:16379"`
	Db       int      `xdefault:"1"`
	Password string   `xdefault:"123456"`
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
