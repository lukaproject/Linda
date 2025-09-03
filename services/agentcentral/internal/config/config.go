package config

type Config struct {
	Env       string `xdefault:"debug" xenv:"env"`
	PGSQL_DSN string `xdefault:"host=localhost user=dxyinme password=123456 dbname=linda port=5432 sslmode=disable TimeZone=Asia/Shanghai" xenv:"PGSQL_DSN"`
	Port      int    `xdefault:"5883"`
	Redis     *RedisConfig
	SSL       *SSLConfig
}

type RedisConfig struct {
	Addrs    []string `xdefault:"host.docker.internal:16379"`
	Db       int      `xdefault:"1"`
	Password string   `xdefault:"123456"`
}

type SSLConfig struct {
	Enabled  bool   `xdefault:"false" xenv:"LINDA_SSL_ENABLED"`
	CertFile string `xdefault:"" xenv:"LINDA_SSL_CERTFILE"`
	KeyFile  string `xdefault:"" xenv:"LINDA_SSL_KEYFILE"`
}

func (c *Config) Merge(other *Config) {
	if other.Env != "" {
		c.Env = other.Env
	}
	if other.PGSQL_DSN != "" {
		c.PGSQL_DSN = other.PGSQL_DSN
	}
	if other.Port != 0 {
		c.Port = other.Port
	}
	if other.Redis != nil {
		c.Redis = other.Redis
	}
	if other.SSL != nil {
		c.SSL = other.SSL
	}
}
