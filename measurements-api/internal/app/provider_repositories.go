package app

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"measurements-api/internal/repository"
	"measurements-api/internal/repository/column"
	"measurements-api/internal/repository/configuration"
	"measurements-api/internal/repository/measurement"
	"measurements-api/internal/repository/object"
	"measurements-api/internal/repository/permission"
	"measurements-api/internal/repository/polling_statistic"
	"measurements-api/internal/repository/post_info"
	"measurements-api/internal/repository/quality"
	"measurements-api/internal/repository/role"
	"measurements-api/internal/repository/session"
	"measurements-api/internal/repository/user"
)

func (p *serviceProvider) ColumnRepository() repository.ColumnRepository {
	if p.columnRepo == nil {
		p.columnRepo = column.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.columnRepo
}

func (p *serviceProvider) ConfigurationRepository() repository.ConfigurationRepository {
	if p.configurationRepo == nil {
		p.configurationRepo = configuration.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.configurationRepo
}

func (p *serviceProvider) MeasurementRepository() repository.MeasurementRepository {
	if p.measurementRepo == nil {
		p.measurementRepo = measurement.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.measurementRepo
}

func (p *serviceProvider) ObjectRepository() repository.ObjectRepository {
	if p.objectRepo == nil {
		p.objectRepo = object.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.objectRepo
}

func (p *serviceProvider) PermissionRepository() repository.PermissionRepository {
	if p.permissionRepo == nil {
		p.permissionRepo = permission.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.permissionRepo
}

func (p *serviceProvider) PollingStatisticRepository() repository.PollingStatisticRepository {
	if p.pollingStatistic == nil {
		p.pollingStatistic = polling_statistic.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.pollingStatistic
}

func (p *serviceProvider) PostInfoRepository() repository.PostInfoRepository {
	if p.postInfoRepo == nil {
		p.postInfoRepo = post_info.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.postInfoRepo
}

func (p *serviceProvider) QualityRepository() repository.QualityRepository {
	if p.qualityRepo == nil {
		p.qualityRepo = quality.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.qualityRepo
}

func (p *serviceProvider) RoleRepository() repository.RoleRepository {
	if p.roleRepo == nil {
		p.roleRepo = role.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.roleRepo
}

func (p *serviceProvider) SessionRepository() repository.SessionRepository {
	if p.sessionRepo == nil {
		p.sessionRepo = session.NewRepository(p.Postgres(), trmpgx.DefaultCtxGetter)
	}
	return p.sessionRepo
}

func (p *serviceProvider) UserRepository() repository.UserRepository {
	if p.userRepo == nil {
		p.userRepo = user.NewRepository(
			p.Postgres(),
			trmpgx.DefaultCtxGetter,
			p.TransactionManager())
	}
	return p.userRepo
}
