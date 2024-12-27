package config

type Config struct {
	Env       string `xdefault:"test" xenv:"env"`
	PGSQL_DSN string `xdefault:"host=localhost user=dxyinme password=123456 dbname=linda port=5432 sslmode=disable TimeZone=Asia/Shanghai"`
	Port      int    `xdefault:"5883"`
	Redis     *RedisConfig
	SSL       *SSLConfig
	FileSaver *FileSaverConfig
}

type RedisConfig struct {
	Addrs    []string `xdefault:"localhost:16379"`
	Db       int      `xdefault:"1"`
	Password string   `xdefault:"123456"`
}

type FileSaverConfig struct {
	RootDir string `xdefault:"/tmp"`
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
	if other.FileSaver != nil {
		c.FileSaver = other.FileSaver
	}
}
