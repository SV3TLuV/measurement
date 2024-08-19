package postgres

import "fmt"

type Config struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     uint16 `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DB       string `env:"POSTGRES_DB"`
}

func (c *Config) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.User, c.Password, c.Addr(), c.DB)
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
