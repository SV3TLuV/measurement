package redis

import "fmt"

type Config struct {
	Host     string `env:"REDIS_HOST"`
	Port     uint16 `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
