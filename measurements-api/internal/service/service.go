package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"measurements-api/internal/model"
	model3 "measurements-api/internal/repository/measurement/model"
	model2 "measurements-api/internal/repository/object/model"
	model5 "measurements-api/internal/repository/polling_statistic/model"
	model4 "measurements-api/internal/repository/user/model"
)

type AppService interface {
	IsOnline() bool
}

type AuthService interface {
	Login(ctx context.Context, request *model.LoginData) (*model.AuthResult, error)
	Logout(ctx context.Context, token *jwt.Token) error
	Refresh(ctx context.Context, token string) (*model.AuthResult, error)
}

type CollectorService interface {
	GetInfo(ctx context.Context) (*model.CollectorInformation, error)
	GetState(ctx context.Context) (*model.CollectorState, error)
	UpdateState(ctx context.Context, state model.CollectorState) error
}

type ColumnService interface {
	Get(ctx context.Context) ([]*model.Column, error)
}

type ConfigurationService interface {
	Get(ctx context.Context) (*model.Configuration, error)
	Update(ctx context.Context, configuration *model.Configuration) error
}

type MeasurementService interface {
	GetMeasurements(
		ctx context.Context,
		options *model3.GetMeasurementsParams) (*model.PagedList[model.Measurement], error)

	GetByID(ctx context.Context, id uint64) (*model.Measurement, error)

	GetLastPostMeasurement(ctx context.Context, postID uint64) (*model.Measurement, error)

	Export(ctx context.Context, options model3.GetMeasurementsParams, fileType model.ExportFormat) ([]byte, error)

	Save(ctx context.Context, measurements []*model.Measurement) (*uint64, error)

	DeleteOutdatedMeasurements(ctx context.Context) error
}

type ObjectService interface {
	GetObjects(
		ctx context.Context,
		options *model2.GetObjectsQueryParams) ([]*model.Object, error)

	GetPosts(ctx context.Context) ([]*model.Object, error)

	GetPost(ctx context.Context, userID, objectID uint64) (*model.Object, error)

	GetTotalPostCount(ctx context.Context) (uint64, error)

	GetListenedPostCount(ctx context.Context) (uint64, error)

	SearchNew(ctx context.Context) ([]*model.Object, error)

	Enable(ctx context.Context, id uint64) error

	Disable(ctx context.Context, id uint64) error
}

type PasswordService interface {
	HashPassword(password string) (*string, error)
	CheckPasswordHash(hashedPassword, password string) bool
}

type PermissionService interface {
	GetPermissions(ctx context.Context) ([]*model.Permission, error)
}

type PollingStatisticService interface {
	GetStatistics(ctx context.Context, options *model5.GetPollingStatisticParams) ([]*model.PollingStatistic, error)
	GetLastStatistic(ctx context.Context) (*model.PollingStatistic, error)
	Create(ctx context.Context, statistic *model.PollingStatistic) error
}

type QualityService interface {
	GetQualities(ctx context.Context) ([]*model.Quality, error)
}

type RoleService interface {
	GetRoles(ctx context.Context) ([]*model.Role, error)
}

type SchedulerService interface {
	Start() error
	Restart() error
}

type TokenService interface {
	CreateAccessToken(claims *model.Claims) (*string, error)
	CreateRefreshToken(claims *model.Claims) (*string, error)
	GetClaims(token string) (*model.Claims, error)
}

type UserService interface {
	GetUsers(ctx context.Context,
		options *model4.GetUsersQueryParams) (*model.PagedList[model.User], error)
	GetUser(ctx context.Context, id uint64) (*model.User, error)
	GetUserObjects(ctx context.Context, userID uint64) ([]*model.Object, error)
	GetUserColumns(ctx context.Context, userID uint64) ([]*model.Column, error)
	GetUserPermissions(ctx context.Context, userID uint64) ([]*model.Permission, error)
	Create(ctx context.Context, request *model.User) error
	Update(ctx context.Context, request *model.User) error
	ChangePassword(ctx context.Context, userID uint64, password string) error
	Ban(ctx context.Context, userID uint64) error
	Unban(ctx context.Context, userID uint64) error
	Delete(ctx context.Context, userID uint64) error
}
