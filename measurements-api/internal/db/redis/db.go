package redis

import "github.com/redis/go-redis/v9"

func NewDB(config *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr(),
		Password: config.Password,
		DB:       config.DB,
	})
}
