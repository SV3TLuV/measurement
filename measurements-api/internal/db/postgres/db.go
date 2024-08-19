package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func NewDB(config *Config) (*pgxpool.Pool, error) {
	ctx := context.Background()
	cfg, err := pgxpool.ParseConfig(config.URL())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse config")
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "create pool")
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "ping")
	}

	return pool, nil
}
