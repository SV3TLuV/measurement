package configuration

import (
	"context"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	def "measurements-api/internal/service"
)

var _ def.ConfigurationService = (*service)(nil)

type service struct {
	repo repository.ConfigurationRepository
}

func NewService(repo repository.ConfigurationRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(ctx context.Context) (*model.Configuration, error) {
	return s.repo.Get(ctx)
}

func (s *service) Update(ctx context.Context, configuration *model.Configuration) error {
	configurationDb, err := s.repo.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get configuration")
	}

	configurationDb.AsoizaLogin = configuration.AsoizaLogin
	configurationDb.AsoizaPassword = configuration.AsoizaPassword
	configurationDb.CollectingInterval = configuration.CollectingInterval
	configurationDb.DeletingInterval = configuration.DeletingInterval
	configurationDb.DeletingThreshold = configuration.DeletingThreshold
	configurationDb.DisablingInterval = configuration.DisablingInterval
	configurationDb.DisablingThreshold = configuration.DisablingThreshold

	err = s.repo.Save(ctx, configurationDb)
	if err != nil {
		return errors.Wrap(err, "failed to update configuration")
	}

	// TODO: restart scheduler

	return nil
}
