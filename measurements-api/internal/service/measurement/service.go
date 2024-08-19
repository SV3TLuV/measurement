package measurement

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	"measurements-api/internal/repository"
	model2 "measurements-api/internal/repository/measurement/model"
	def "measurements-api/internal/service"
	exporter2 "measurements-api/pkg/exporter"
	"reflect"
	"time"
)

var _ def.MeasurementService = (*service)(nil)

type service struct {
	repo                 repository.MeasurementRepository
	configurationService def.ConfigurationService
	userService          def.UserService
}

func NewService(repo repository.MeasurementRepository,
	configurationService def.ConfigurationService,
	userService def.UserService) *service {
	return &service{
		repo:                 repo,
		configurationService: configurationService,
		userService:          userService,
	}
}

func (u *service) GetMeasurements(
	ctx context.Context,
	options *model2.GetMeasurementsParams) (*model.PagedList[model.Measurement], error) {
	measurements, count, err := u.repo.Get(ctx, options)
	if err != nil {
		return nil, errors.Wrap(err, "get users")
	}

	pagedList := &model.PagedList[model.Measurement]{
		Page:     options.Page,
		PageSize: options.PageSize,
		Total:    *count,
		Items:    measurements,
	}

	return pagedList, nil
}

func (u *service) GetByID(ctx context.Context, id uint64) (*model.Measurement, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *service) GetLastPostMeasurement(ctx context.Context, postID uint64) (*model.Measurement, error) {
	return u.repo.GetLastPostMeasurement(ctx, postID)
}

func (u *service) Export(ctx context.Context, options model2.GetMeasurementsParams, fileType model.ExportFormat) ([]byte, error) {
	columns, err := u.userService.GetUserColumns(ctx, options.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "get user columns")
	}

	measurements, _, err := u.repo.Get(ctx, &options)
	if err != nil {
		return nil, errors.Wrap(err, "get measurements")
	}

	data := make([][]string, len(measurements)+1)
	for i := 0; i < cap(data); i++ {
		data[i] = make([]string, 0, len(columns))
	}

	objFields := make([]string, 0, len(columns))
	for _, column := range columns {
		data[0] = append(data[0], column.Title)
		objFields = append(objFields, column.GetFormattedObjectField())
	}

	for i := 0; i < len(measurements); i++ {
		val := reflect.ValueOf(measurements[i])
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				data[i+1] = append(data[i+1], "")
				continue
			}
			val = val.Elem()
		}
		for j := 0; j < len(objFields); j++ {
			field := val.FieldByName(objFields[j])
			if field.IsValid() && field.CanInterface() {
				data[i+1] = append(data[i+1], fmt.Sprintf("%v", field.Interface()))
			} else {
				data[i+1] = append(data[i+1], "")
			}
		}
	}

	exporter := exporter2.NewExporter(fileType)
	if exporter == nil {
		return nil, errors.New("exporter is nil")
	}

	return exporter.Export(data)
}

func (u *service) Save(ctx context.Context, measurements []*model.Measurement) (*uint64, error) {
	return u.repo.Save(ctx, measurements)
}

func (u *service) DeleteOutdatedMeasurements(ctx context.Context) error {
	configuration, err := u.configurationService.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "get configuration")
	}

	month := configuration.DeletingThreshold / 2592000 // convert to months
	threshold := time.Now().UTC().AddDate(0, -int(month), 0)
	err = u.repo.DeleteCreatedBefore(ctx, &threshold)
	if err != nil {
		return errors.Wrap(err, "delete created before")
	}

	return nil
}
