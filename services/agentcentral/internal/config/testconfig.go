package config

func TestConfig() *Config {
	return &Config{
		PGSQL_DSN: "host=localhost user=dxyinme password=123456 dbname=linda_test port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		Port:      5883,
	}
}
