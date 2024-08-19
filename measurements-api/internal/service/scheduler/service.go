package scheduler

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"measurements-api/internal/repository"
	"measurements-api/internal/server/scheduler/jobs"
	service2 "measurements-api/internal/service"
	"measurements-api/pkg/asoiza"
)

type service struct {
	cron                 *cron.Cron
	client               asoiza.Client
	measurementService   service2.MeasurementService
	configurationService service2.ConfigurationService
	postInfoRepo         repository.PostInfoRepository
	objectService        service2.ObjectService
	collectorService     service2.CollectorService
}

func NewService(
	cron *cron.Cron,
	client asoiza.Client,
	measurementService service2.MeasurementService,
	configurationService service2.ConfigurationService,
	postInfoRepo repository.PostInfoRepository,
	objectService service2.ObjectService,
	collectorService service2.CollectorService) *service {
	return &service{
		cron:                 cron,
		client:               client,
		measurementService:   measurementService,
		configurationService: configurationService,
		postInfoRepo:         postInfoRepo,
		objectService:        objectService,
		collectorService:     collectorService,
	}
}

func (s *service) Start() error {
	err := s.initJobs()
	if err != nil {
		return errors.Wrap(err, "init jobs")
	}
	s.cron.Start()
	return nil
}

func (s *service) Restart() error {
	s.cron.Stop()
	for _, task := range s.cron.Entries() {
		s.cron.Remove(task.ID)
	}
	return s.Start()
}

func (s *service) initJobs() error {
	configuration, err := s.configurationService.Get(context.Background())
	if err != nil {
		return errors.Wrap(err, "get configuration")
	}

	collectingMinutes := configuration.CollectingInterval / 60
	collectingCron := fmt.Sprintf("0 3/%d * * * *", collectingMinutes)
	_, err = s.cron.AddFunc(collectingCron, jobs.NewCollectMeasurementJob(
		s.objectService, s.client, s.postInfoRepo, s.measurementService, s.collectorService))
	if err != nil {
		return errors.Wrap(err, "add NewCollectMeasurementJob")
	}

	disablingHours := configuration.DisablingInterval / 3600
	disablingCron := fmt.Sprintf("0 0 */%d * * *", disablingHours)
	_, err = s.cron.AddFunc(disablingCron, jobs.NewDisablingPostsJob(
		s.configurationService, s.measurementService, s.objectService, s.client))
	if err != nil {
		return errors.Wrap(err, "add NewDisablingPostsJob")
	}

	deletingDays := configuration.DeletingInterval / (3600 * 24)
	deletingCron := fmt.Sprintf("0 0 0 1/%d * *", deletingDays)
	_, err = s.cron.AddFunc(deletingCron, jobs.NewDeletingOutdatedMeasurementJob(
		s.measurementService))
	if err != nil {
		return errors.Wrap(err, "add NewDeletingOutdatedMeasurementJob")
	}

	return nil
}
