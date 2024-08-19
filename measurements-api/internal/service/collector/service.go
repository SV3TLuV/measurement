package collector

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	redis2 "measurements-api/internal/db/redis"
	"measurements-api/internal/model"
	service2 "measurements-api/internal/service"
	"time"
)

var _ service2.CollectorService = (*service)(nil)

type service struct {
	statisticService     service2.PollingStatisticService
	configurationService service2.ConfigurationService
	objectService        service2.ObjectService
	redis                *redis.Client
}

func NewService(
	statisticService service2.PollingStatisticService,
	configurationService service2.ConfigurationService,
	objectService service2.ObjectService,
	redis *redis.Client) *service {
	return &service{
		statisticService:     statisticService,
		configurationService: configurationService,
		objectService:        objectService,
		redis:                redis,
	}
}

func (s *service) GetInfo(ctx context.Context) (*model.CollectorInformation, error) {
	configuration, err := s.configurationService.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get configuration")
	}

	information := &model.CollectorInformation{
		PollingInterval: configuration.CollectingInterval,
	}

	information.PostCount, err = s.objectService.GetTotalPostCount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get total")
	}

	information.ListenedPostCount, err = s.objectService.GetListenedPostCount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get listened post count")
	}

	statistic, err := s.statisticService.GetLastStatistic(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get statistic")
	}
	if statistic != nil {
		interval := time.Duration(configuration.CollectingInterval) * time.Second
		untilNextPolling := statistic.DateTime.Add(interval)
		information.LastPollingDateTime = &statistic.DateTime
		information.UntilNextPolling = &untilNextPolling
	}

	return information, nil
}

func (s *service) GetState(ctx context.Context) (*model.CollectorState, error) {
	state, err := redis2.Get[model.CollectorState](s.redis, ctx, "collector:state")
	if errors.Is(err, redis.Nil) {
		err = s.ResetState(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "init state")
		}

		state, err = redis2.Get[model.CollectorState](s.redis, ctx, "collector:state")
	}
	if err != nil {
		return nil, errors.Wrap(err, "get state")
	}

	return state, nil
}

func (s *service) UpdateState(ctx context.Context, state model.CollectorState) error {
	if state.Ended != nil {
		defer s.ResetState(ctx)

		err := s.statisticService.Create(ctx, &model.PollingStatistic{
			DateTime:      *state.Started,
			Duration:      state.Ended.Sub(*state.Started),
			PostCount:     state.PostCount,
			ReceivedCount: state.ReceivedCount,
		})
		if err != nil {
			return errors.Wrap(err, "update state")
		}

		return nil
	}

	err := redis2.Set(s.redis, ctx, "collector:state", state, redis.KeepTTL)
	if err != nil {
		return errors.Wrap(err, "update state")
	}

	return nil
}

func (s *service) ResetState(ctx context.Context) error {
	return s.UpdateState(ctx, model.CollectorState{
		Status:          model.Pending,
		PolledPostCount: 0,
		PostCount:       0,
		PollingPercent:  0,
		ReceivedCount:   0,
		Started:         nil,
		Ended:           nil,
	})
}
