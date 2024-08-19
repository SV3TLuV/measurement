package app

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"measurements-api/internal/config"
	"measurements-api/internal/db/postgres"
	redis2 "measurements-api/internal/db/redis"
	asoiza2 "measurements-api/internal/interfaces/factories/asoiza"
	"measurements-api/internal/repository"
	"measurements-api/internal/service"
	"measurements-api/pkg/asoiza"
)

type serviceProvider struct {
	config *config.Config

	cron *cron.Cron

	postgresPool *pgxpool.Pool
	trManager    *manager.Manager

	redis *redis.Client

	asoizaClient asoiza.Client

	appService  service.AppService
	authService service.AuthService

	collectorService service.CollectorService

	columnRepo    repository.ColumnRepository
	columnService service.ColumnService

	configurationRepo    repository.ConfigurationRepository
	configurationService service.ConfigurationService

	measurementRepo    repository.MeasurementRepository
	measurementService service.MeasurementService

	objectRepo    repository.ObjectRepository
	objectService service.ObjectService

	permissionRepo    repository.PermissionRepository
	permissionService service.PermissionService

	pollingStatistic        repository.PollingStatisticRepository
	pollingStatisticService service.PollingStatisticService

	postInfoRepo repository.PostInfoRepository

	qualityRepo    repository.QualityRepository
	qualityService service.QualityService

	roleRepo    repository.RoleRepository
	roleService service.RoleService

	schedulerService service.SchedulerService

	sessionRepo repository.SessionRepository

	userRepo    repository.UserRepository
	userService service.UserService

	passwordService service.PasswordService
	tokenService    service.TokenService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (p *serviceProvider) Config() *config.Config {
	if p.config == nil {
		p.config = config.FromEnv()
	}
	return p.config
}

func (p *serviceProvider) Postgres() *pgxpool.Pool {
	if p.postgresPool == nil {
		db, err := postgres.NewDB(p.Config().Postgres)
		if err != nil {
			panic(err)
		}
		p.postgresPool = db
	}
	return p.postgresPool
}

func (p *serviceProvider) Redis() *redis.Client {
	if p.redis == nil {
		p.redis = redis2.NewDB(p.Config().Redis)
	}
	return p.redis
}

func (p *serviceProvider) TransactionManager() *manager.Manager {
	if p.trManager == nil {
		p.trManager = manager.Must(trmpgx.NewDefaultFactory(p.postgresPool))
	}
	return p.trManager
}

func (p *serviceProvider) AsoizaClient() asoiza.Client {
	if p.asoizaClient == nil {
		p.asoizaClient = asoiza.NewClient(
			asoiza2.NewConfigurationFactory(p.ConfigurationService()))
	}
	return p.asoizaClient
}

func (p *serviceProvider) Cron() *cron.Cron {
	if p.cron == nil {
		p.cron = cron.New(cron.WithSeconds())
	}
	return p.cron
}
