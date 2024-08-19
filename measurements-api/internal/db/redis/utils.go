package redis

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func Get[T any](client *redis.Client, ctx context.Context, key string) (*T, error) {
	jsonValue, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}

	var value T
	err = json.Unmarshal([]byte(jsonValue), &value)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	return &value, nil
}

func Set(client *redis.Client, ctx context.Context, key string, value any, duration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	err = client.Set(ctx, key, string(jsonValue), duration).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set")
	}

	return nil
}
