package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"measurements-api/internal/db/postgres"
	"measurements-api/internal/db/redis"
)

type Config struct {
	Http     *HttpConfig      `env:""`
	Postgres *postgres.Config `env:""`
	Redis    *redis.Config    `env:""`
	Jwt      *JwtConfig       `env:""`
}

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("Error loading .env file")
	}
	return nil
}

func FromEnv() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		panic(err)
	}

	return cfg
}
