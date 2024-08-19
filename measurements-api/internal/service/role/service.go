package role

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	redis2 "measurements-api/internal/db/redis"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	def "measurements-api/internal/service"
)

var _ def.RoleService = (*service)(nil)

type service struct {
	repo  repository.RoleRepository
	redis *redis.Client
}

func NewService(repo repository.RoleRepository,
	redis *redis.Client) *service {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

func (s *service) GetRoles(ctx context.Context) ([]*model.Role, error) {
	const key = "roles:all"

	cache, err := redis2.Get[[]*model.Role](s.redis, ctx, key)
	if errors.Is(err, redis.Nil) {
		roles, err := s.repo.Get(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get roles")
		}

		err = redis2.Set(s.redis, ctx, key, roles, redis.KeepTTL)
		if err != nil {
			return nil, errors.Wrap(err, "cache roles")
		}

		return roles, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "get roles cache")
	}

	return *cache, nil
}
