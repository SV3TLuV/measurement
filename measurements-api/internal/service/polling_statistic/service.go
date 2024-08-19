package polling_statistic

import (
	"context"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	model2 "measurements-api/internal/repository/polling_statistic/model"
	def "measurements-api/internal/service"
)

var _ def.PollingStatisticService = (*service)(nil)

type service struct {
	repo repository.PollingStatisticRepository
}

func NewService(repo repository.PollingStatisticRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetStatistics(ctx context.Context, options *model2.GetPollingStatisticParams) ([]*model.PollingStatistic, error) {
	return s.repo.Get(ctx, options)
}

func (s *service) GetLastStatistic(ctx context.Context) (*model.PollingStatistic, error) {
	statistics, err := s.repo.Get(ctx, &model2.GetPollingStatisticParams{
		Page:     1,
		PageSize: 1,
	})
	if err != nil {
		return nil, err
	}
	if len(statistics) == 0 {
		return nil, nil
	}
	return statistics[0], nil
}

func (s *service) Create(ctx context.Context, statistic *model.PollingStatistic) error {
	return s.repo.SaveOne(ctx, statistic)
}
