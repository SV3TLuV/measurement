package column

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	redis2 "measurements-api/internal/db/redis"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	def "measurements-api/internal/service"
)

var _ def.ColumnService = (*service)(nil)

type service struct {
	repo  repository.ColumnRepository
	redis *redis.Client
}

func NewService(repo repository.ColumnRepository,
	redis *redis.Client) *service {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

func (s *service) Get(ctx context.Context) ([]*model.Column, error) {
	const key = "columns:all"

	cache, err := redis2.Get[[]*model.Column](s.redis, ctx, key)
	if errors.Is(err, redis.Nil) {
		columns, err := s.repo.Get(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get columns")
		}

		err = redis2.Set(s.redis, ctx, key, columns, redis.KeepTTL)
		if err != nil {
			return nil, errors.Wrap(err, "cache columns")
		}

		return columns, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "get columns cache")
	}

	return *cache, nil
}
