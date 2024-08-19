package permission

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	redis2 "measurements-api/internal/db/redis"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	def "measurements-api/internal/service"
)

var _ def.PermissionService = (*service)(nil)

type service struct {
	repo  repository.PermissionRepository
	redis *redis.Client
}

func NewService(repo repository.PermissionRepository,
	redis *redis.Client) *service {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

func (s *service) GetPermissions(ctx context.Context) ([]*model.Permission, error) {
	const key = "permissions:all"

	cache, err := redis2.Get[[]*model.Permission](s.redis, ctx, key)
	if errors.Is(err, redis.Nil) {
		permissions, err := s.repo.Get(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get permissions")
		}

		err = redis2.Set(s.redis, ctx, key, permissions, redis.KeepTTL)
		if err != nil {
			return nil, errors.Wrap(err, "cache permissions")
		}

		return permissions, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "get permissions cache")
	}

	return *cache, nil
}
