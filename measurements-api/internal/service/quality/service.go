package quality

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	redis2 "measurements-api/internal/db/redis"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	def "measurements-api/internal/service"
)

var _ def.QualityService = (*service)(nil)

type service struct {
	repo  repository.QualityRepository
	redis *redis.Client
}

func NewService(repo repository.QualityRepository,
	redis *redis.Client) *service {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

func (s *service) GetQualities(ctx context.Context) ([]*model.Quality, error) {
	const key = "qualities:all"

	cache, err := redis2.Get[[]*model.Quality](s.redis, ctx, key)
	if errors.Is(err, redis.Nil) {
		qualities, err := s.repo.Get(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get qualities")
		}

		err = redis2.Set(s.redis, ctx, key, qualities, redis.KeepTTL)
		if err != nil {
			return nil, errors.Wrap(err, "cache qualities")
		}

		return qualities, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "get qualities cache")
	}

	return *cache, nil
}
