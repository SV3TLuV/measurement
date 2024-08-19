package app

import (
	"measurements-api/internal/service"
	"measurements-api/internal/service/app"
	"measurements-api/internal/service/auth"
	"measurements-api/internal/service/collector"
	column2 "measurements-api/internal/service/column"
	"measurements-api/internal/service/configuration"
	"measurements-api/internal/service/measurement"
	"measurements-api/internal/service/object"
	"measurements-api/internal/service/password"
	"measurements-api/internal/service/permission"
	"measurements-api/internal/service/polling_statistic"
	"measurements-api/internal/service/quality"
	"measurements-api/internal/service/role"
	"measurements-api/internal/service/scheduler"
	"measurements-api/internal/service/token"
	"measurements-api/internal/service/user"
)

func (p *serviceProvider) AppService() service.AppService {
	if p.appService == nil {
		p.appService = app.NewService()
	}
	return p.appService
}

func (p *serviceProvider) AuthService() service.AuthService {
	if p.authService == nil {
		p.authService = auth.NewService(
			p.UserRepository(),
			p.SessionRepository(),
			p.TokenService(),
			p.PasswordService(),
		)
	}
	return p.authService
}

func (p *serviceProvider) CollectorService() service.CollectorService {
	if p.collectorService == nil {
		p.collectorService = collector.NewService(
			p.PollingStatisticService(),
			p.ConfigurationService(),
			p.ObjectService(),
			p.Redis())
	}
	return p.collectorService
}

func (p *serviceProvider) ColumnService() service.ColumnService {
	if p.columnService == nil {
		p.columnService = column2.NewService(
			p.ColumnRepository(),
			p.Redis(),
		)
	}
	return p.columnService
}

func (p *serviceProvider) ConfigurationService() service.ConfigurationService {
	if p.configurationService == nil {
		p.configurationService = configuration.NewService(
			p.ConfigurationRepository(),
		)
	}
	return p.configurationService
}

func (p *serviceProvider) MeasurementService() service.MeasurementService {
	if p.measurementService == nil {
		p.measurementService = measurement.NewService(
			p.MeasurementRepository(),
			p.ConfigurationService(),
			p.UserService(),
		)
	}
	return p.measurementService
}

func (p *serviceProvider) ObjectService() service.ObjectService {
	if p.objectService == nil {
		p.objectService = object.NewService(
			p.ObjectRepository(),
			p.PostInfoRepository(),
			p.AsoizaClient(),
		)
	}
	return p.objectService
}

func (p *serviceProvider) PasswordService() service.PasswordService {
	if p.passwordService == nil {
		p.passwordService = password.NewService()
	}
	return p.passwordService
}

func (p *serviceProvider) PermissionService() service.PermissionService {
	if p.permissionService == nil {
		p.permissionService = permission.NewService(
			p.PermissionRepository(),
			p.Redis(),
		)
	}
	return p.permissionService
}

func (p *serviceProvider) PollingStatisticService() service.PollingStatisticService {
	if p.pollingStatisticService == nil {
		p.pollingStatisticService = polling_statistic.NewService(p.PollingStatisticRepository())
	}
	return p.pollingStatisticService
}

func (p *serviceProvider) QualityService() service.QualityService {
	if p.qualityService == nil {
		p.qualityService = quality.NewService(
			p.QualityRepository(),
			p.Redis(),
		)
	}
	return p.qualityService
}

func (p *serviceProvider) RoleService() service.RoleService {
	if p.roleService == nil {
		p.roleService = role.NewService(
			p.RoleRepository(),
			p.Redis(),
		)
	}
	return p.roleService
}

func (p *serviceProvider) SchedulerService() service.SchedulerService {
	if p.schedulerService == nil {
		p.schedulerService = scheduler.NewService(
			p.Cron(),
			p.AsoizaClient(),
			p.MeasurementService(),
			p.ConfigurationService(),
			p.PostInfoRepository(),
			p.ObjectService(),
			p.CollectorService())
	}
	return p.schedulerService
}

func (p *serviceProvider) TokenService() service.TokenService {
	if p.tokenService == nil {
		p.tokenService = token.NewService(p.Config().Jwt)
	}
	return p.tokenService
}

func (p *serviceProvider) UserService() service.UserService {
	if p.userService == nil {
		p.userService = user.NewService(
			p.UserRepository(),
			p.PasswordService(),
			p.TransactionManager(),
		)
	}
	return p.userService
}
