package config

import "fmt"

type HttpConfig struct {
	Host string `env:"HTTP_HOST"`
	Port uint16 `env:"HTTP_PORT"`
}

func NewConfig(host string, port uint16) *HttpConfig {
	return &HttpConfig{
		Host: host,
		Port: port,
	}
}

func (c *HttpConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
