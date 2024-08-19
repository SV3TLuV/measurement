package repository

import (
	"context"
	"github.com/google/uuid"
	"measurements-api/internal/model"
	model3 "measurements-api/internal/repository/measurement/model"
	model2 "measurements-api/internal/repository/object/model"
	model5 "measurements-api/internal/repository/polling_statistic/model"
	model4 "measurements-api/internal/repository/user/model"
	"time"
)

type ColumnRepository interface {
	Get(ctx context.Context) ([]*model.Column, error)
}

type ConfigurationRepository interface {
	Get(ctx context.Context) (*model.Configuration, error)
	Save(ctx context.Context, configuration *model.Configuration) error
}

type MeasurementRepository interface {
	Get(ctx context.Context, options *model3.GetMeasurementsParams) ([]*model.Measurement, *uint64, error)
	GetByID(ctx context.Context, id uint64) (*model.Measurement, error)
	GetLastPostMeasurement(ctx context.Context, postID uint64) (*model.Measurement, error)
	Save(ctx context.Context, measurements []*model.Measurement) (*uint64, error)
	SaveOne(ctx context.Context, measurement *model.Measurement) error
	DeleteCreatedBefore(ctx context.Context, before *time.Time) error
}

type ObjectRepository interface {
	Get(ctx context.Context, options *model2.GetObjectsQueryParams) ([]*model.Object, error)
	GetByIds(ctx context.Context, ids []uint64) ([]*model.Object, error)
	GetById(ctx context.Context, id uint64) (*model.Object, error)
	GetUserPostById(ctx context.Context, userID, postID uint64) (*model.Object, error)
	GetCount(ctx context.Context, options *model2.GetObjectCountParams) (uint64, error)
	Save(ctx context.Context, objects []*model.Object) ([]*model.Object, error)
}

type PermissionRepository interface {
	Get(ctx context.Context) ([]*model.Permission, error)
}

type PollingStatisticRepository interface {
	Get(ctx context.Context, options *model5.GetPollingStatisticParams) ([]*model.PollingStatistic, error)
	SaveOne(ctx context.Context, statistic *model.PollingStatistic) error
}

type PostInfoRepository interface {
	GetById(ctx context.Context, id uint64) (*model.PostInfo, error)
	Save(ctx context.Context, infos []*model.PostInfo) error
	SaveOne(ctx context.Context, infos *model.PostInfo) error
}

type QualityRepository interface {
	Get(ctx context.Context) ([]*model.Quality, error)
}

type RoleRepository interface {
	Get(ctx context.Context) ([]*model.Role, error)
}

type SessionRepository interface {
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]*model.Session, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.Session, error)
	Save(ctx context.Context, session *model.Session) error
	Delete(ctx context.Context, ids []uuid.UUID) error
	DeleteOne(ctx context.Context, id uuid.UUID) error
}

type UserRepository interface {
	Get(ctx context.Context,
		options *model4.GetUsersQueryParams) (users []*model.User, total *uint64, err error)
	GetByIds(ctx context.Context, ids []uint64) ([]*model.User, error)
	GetById(ctx context.Context, id uint64) (*model.User, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
	Save(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
	GetUserPermissions(ctx context.Context, userID uint64) ([]*model.Permission, error)
	UpdateUserPermissions(ctx context.Context, userID uint64, permissionIDs []uint64) error
	GetUserColumns(ctx context.Context, userID uint64) ([]*model.Column, error)
	UpdateUserColumns(ctx context.Context, userID uint64, columnIDs []uint64) error
	GetUserObjects(ctx context.Context, userID uint64) ([]*model.Object, error)
	UpdateUserPosts(ctx context.Context, userID uint64, postIDs []uint64) error
}
