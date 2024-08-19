package measurement

import (
	"context"
	"fmt"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	def "measurements-api/internal/repository"
	"measurements-api/internal/repository/entities"
	"measurements-api/internal/repository/measurement/converter"
	model2 "measurements-api/internal/repository/measurement/model"
	"strings"
	"time"
)

var _ def.MeasurementRepository = (*repository)(nil)

type repository struct {
	pool   *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewRepository(pool *pgxpool.Pool, c *trmpgx.CtxGetter) *repository {
	return &repository{
		pool:   pool,
		getter: c,
	}
}

func (r *repository) Get(
	ctx context.Context,
	options *model2.GetMeasurementsParams) ([]*model.Measurement, *uint64, error) {
	if options == nil {
		return nil, nil, errors.Wrap(model.Empty, "options")
	}

	dialect := goqu.Dialect("postgres")
	fieldsQuery := dialect.From("user_columns").
		Select(goqu.L("array_agg(columns.obj_field)")).
		Join(
			goqu.T("columns"),
			goqu.On(goqu.Ex{"columns.column_id": goqu.I("user_columns.column_id")})).
		Where(goqu.Ex{"user_columns.user_id": options.UserID})

	query := dialect.From("measurements").
		Join(goqu.T("user_posts"), goqu.On(goqu.Ex{
			"user_posts.object_id": goqu.I("measurements.object_id"),
		})).
		Where(goqu.Ex{
			"user_posts.user_id": options.UserID,
		})

	if options.ObjectID != nil {
		query = query.Where(goqu.Ex{
			"measurements.object_id": *options.ObjectID,
		})
	}
	if options.Start != nil {
		start := options.Start.Format("2006-01-02")
		query = query.Where(goqu.L("measurements.real_date_time::date >=", start))
	}
	if options.End != nil {
		end := options.End.Format("2006-01-02")
		query = query.Where(goqu.L("measurements.real_date_time::date <=", end))
	}
	if options.Period != nil && *options.Period != model.All {
		now := time.Now()
		duration := options.Period.ToDuration()
		period := now.Add(-duration).Format("2006-01-02")
		query = query.Where(goqu.L("measurements.real_date_time::date >=", period))
	}

	countQuery := query.Clone().(*goqu.SelectDataset).
		Select(goqu.L("COUNT(*)").As("total_count"))

	selectQuery := query.Clone().(*goqu.SelectDataset).
		Select(
			goqu.I("measurements.measurement_id"),
			goqu.I("measurements.object_id"),
			goqu.I("measurements.created"),
			goqu.I("measurements.changed"),
		).
		Order(goqu.I("measurement_id").Desc())

	if options.Page > 0 && options.PageSize > 0 {
		selectQuery = selectQuery.Limit(options.PageSize).
			Offset((options.Page - 1) * options.PageSize)
	}

	var columns []string
	sql, args, err := fieldsQuery.ToSQL()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate query")
	}
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &columns, sql, args...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "select")
	}

	for i := 0; i < len(columns); i++ {
		if strings.HasPrefix(columns[i], "v_") {
			columns = append(columns,
				strings.Replace(columns[i], "v_", "q_", 1),
				strings.Replace(columns[i], "v_", "m_", 1))
		}
	}

	for i := 0; i < len(columns); i++ {
		selectQuery = selectQuery.SelectAppend(goqu.I(fmt.Sprintf("measurements.%s", columns[i])))
	}

	var measurements []*entities.Measurement
	sql, args, err = selectQuery.ToSQL()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate query")
	}
	err = pgxscan.Select(ctx, tr, &measurements, sql, args...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "select")
	}

	sql, args, err = countQuery.ToSQL()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate count sql")
	}

	var total uint64
	err = tr.QueryRow(ctx, sql, args...).Scan(&total)
	if err != nil {
		return nil, nil, errors.Wrap(err, "scan total")
	}

	return converter.ToMeasurementsFromRepo(measurements), &total, nil
}

func (r *repository) GetLastPostMeasurement(ctx context.Context, postID uint64) (*model.Measurement, error) {
	query := goqu.Dialect("postgres").
		From("measurements").
		Order(goqu.I("measurement_id").Desc()).
		Where(goqu.Ex{"object_id": postID}).
		Limit(1)

	var measurement entities.Measurement
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &measurement, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToMeasurementFromRepo(&measurement), nil
}

func (r *repository) GetByID(ctx context.Context, id uint64) (*model.Measurement, error) {
	query := goqu.Dialect("postgres").
		From("measurements").
		Where(goqu.Ex{"measurement_id": id})

	var measurement entities.Measurement
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &measurement, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToMeasurementFromRepo(&measurement), nil
}

func (r *repository) Save(ctx context.Context, measurements []*model.Measurement) (*uint64, error) {
	if len(measurements) == 0 {
		return nil, nil
	}

	records := converter.ToMeasurementRecordsFromService(measurements)
	query := goqu.Dialect("postgres").
		Insert("measurements").
		Rows(records).
		OnConflict(
			goqu.DoUpdate("measurement_id", &goqu.Record{
				"obj_name":               goqu.L("EXCLUDED.obj_name"),
				"obj_num":                goqu.L("EXCLUDED.obj_num"),
				"created":                goqu.L("EXCLUDED.created"),
				"changed":                goqu.L("EXCLUDED.changed"),
				"date_time":              goqu.L("EXCLUDED.date_time"),
				"real_date_time":         goqu.L("EXCLUDED.real_date_time"),
				"temp":                   goqu.L("EXCLUDED.temp"),
				"pressure":               goqu.L("EXCLUDED.pressure"),
				"wind_dir":               goqu.L("EXCLUDED.wind_dir"),
				"wind_dir_str":           goqu.L("EXCLUDED.wind_dir_str"),
				"wind_speed":             goqu.L("EXCLUDED.wind_speed"),
				"humid":                  goqu.L("EXCLUDED.humid"),
				"water_vapor_elasticity": goqu.L("EXCLUDED.water_vapor_elasticity"),
				"atm_phenom":             goqu.L("EXCLUDED.atm_phenom"),
				"humid_int":              goqu.L("EXCLUDED.humid_int"),
				"temp_int":               goqu.L("EXCLUDED.temp_int"),
				"v_202917":               goqu.L("EXCLUDED.v_202917"),
				"m_202917":               goqu.L("EXCLUDED.m_202917"),
				"q_202917":               goqu.L("EXCLUDED.q_202917"),
				"v_202918":               goqu.L("EXCLUDED.v_202918"),
				"m_202918":               goqu.L("EXCLUDED.m_202918"),
				"q_202918":               goqu.L("EXCLUDED.q_202918"),
				"v_202919":               goqu.L("EXCLUDED.v_202919"),
				"m_202919":               goqu.L("EXCLUDED.m_202919"),
				"q_202919":               goqu.L("EXCLUDED.q_202919"),
				"v_202920":               goqu.L("EXCLUDED.v_202920"),
				"m_202920":               goqu.L("EXCLUDED.m_202920"),
				"q_202920":               goqu.L("EXCLUDED.q_202920"),
				"v_202921":               goqu.L("EXCLUDED.v_202921"),
				"m_202921":               goqu.L("EXCLUDED.m_202921"),
				"q_202921":               goqu.L("EXCLUDED.q_202921"),
				"v_202932":               goqu.L("EXCLUDED.v_202932"),
				"m_202932":               goqu.L("EXCLUDED.m_202932"),
				"q_202932":               goqu.L("EXCLUDED.q_202932"),
				"v_202935":               goqu.L("EXCLUDED.v_202935"),
				"m_202935":               goqu.L("EXCLUDED.m_202935"),
				"q_202935":               goqu.L("EXCLUDED.q_202935"),
				"v_202924":               goqu.L("EXCLUDED.v_202924"),
				"m_202924":               goqu.L("EXCLUDED.m_202924"),
				"q_202924":               goqu.L("EXCLUDED.q_202924"),
				"v_202925":               goqu.L("EXCLUDED.v_202925"),
				"m_202925":               goqu.L("EXCLUDED.m_202925"),
				"q_202925":               goqu.L("EXCLUDED.q_202925"),
				"v_203565":               goqu.L("EXCLUDED.v_203565"),
				"m_203565":               goqu.L("EXCLUDED.m_203565"),
				"q_203565":               goqu.L("EXCLUDED.q_203565"),
				"v_209190":               goqu.L("EXCLUDED.v_209190"),
				"m_209190":               goqu.L("EXCLUDED.m_209190"),
				"q_209190":               goqu.L("EXCLUDED.q_209190"),
				"v_203570":               goqu.L("EXCLUDED.v_203570"),
				"m_203570":               goqu.L("EXCLUDED.m_203570"),
				"q_203570":               goqu.L("EXCLUDED.q_203570"),
				"v_203551":               goqu.L("EXCLUDED.v_203551"),
				"m_203551":               goqu.L("EXCLUDED.m_203551"),
				"q_203551":               goqu.L("EXCLUDED.q_203551"),
				"v_202936":               goqu.L("EXCLUDED.v_202936"),
				"m_202936":               goqu.L("EXCLUDED.m_202936"),
				"q_202936":               goqu.L("EXCLUDED.q_202936"),
				"v_203569":               goqu.L("EXCLUDED.v_203569"),
				"m_203569":               goqu.L("EXCLUDED.m_203569"),
				"q_203569":               goqu.L("EXCLUDED.q_203569"),
				"v_203557":               goqu.L("EXCLUDED.v_203557"),
				"m_203557":               goqu.L("EXCLUDED.m_203557"),
				"q_203557":               goqu.L("EXCLUDED.q_203557"),
				"v_203568":               goqu.L("EXCLUDED.v_203568"),
				"m_203568":               goqu.L("EXCLUDED.m_203568"),
				"q_203568":               goqu.L("EXCLUDED.q_203568"),
				"v_203559":               goqu.L("EXCLUDED.v_203559"),
				"m_203559":               goqu.L("EXCLUDED.m_203559"),
				"q_203559":               goqu.L("EXCLUDED.q_203559"),
				"v_203577":               goqu.L("EXCLUDED.v_203577"),
				"m_203577":               goqu.L("EXCLUDED.m_203577"),
				"q_203577":               goqu.L("EXCLUDED.q_203577"),
				"v_211082":               goqu.L("EXCLUDED.v_211082"),
				"m_211082":               goqu.L("EXCLUDED.m_211082"),
				"q_211082":               goqu.L("EXCLUDED.q_211082"),
				"v_202931":               goqu.L("EXCLUDED.v_202931"),
				"m_202931":               goqu.L("EXCLUDED.m_202931"),
				"q_202931":               goqu.L("EXCLUDED.q_202931"),
			}),
		).
		Returning("measurement_id")
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate query")
	}

	var insertedIDs []uint64
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &insertedIDs, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "save measurements")
	}

	insertedCount := uint64(len(insertedIDs))
	return &insertedCount, nil
}

func (r *repository) SaveOne(ctx context.Context, measurement *model.Measurement) error {
	record := converter.ToMeasurementRecordFromService(measurement)
	query := goqu.Dialect("postgres").
		Insert("measurements").
		Rows(record).
		OnConflict(
			goqu.DoUpdate("measurement_id", &goqu.Record{
				"obj_name":               goqu.L("EXCLUDED.obj_name"),
				"obj_num":                goqu.L("EXCLUDED.obj_num"),
				"created":                goqu.L("EXCLUDED.created"),
				"changed":                goqu.L("EXCLUDED.changed"),
				"date_time":              goqu.L("EXCLUDED.date_time"),
				"real_date_time":         goqu.L("EXCLUDED.real_date_time"),
				"temp":                   goqu.L("EXCLUDED.temp"),
				"pressure":               goqu.L("EXCLUDED.pressure"),
				"wind_dir":               goqu.L("EXCLUDED.wind_dir"),
				"wind_dir_str":           goqu.L("EXCLUDED.wind_dir_str"),
				"wind_speed":             goqu.L("EXCLUDED.wind_speed"),
				"humid":                  goqu.L("EXCLUDED.humid"),
				"water_vapor_elasticity": goqu.L("EXCLUDED.water_vapor_elasticity"),
				"atm_phenom":             goqu.L("EXCLUDED.atm_phenom"),
				"humid_int":              goqu.L("EXCLUDED.humid_int"),
				"temp_int":               goqu.L("EXCLUDED.temp_int"),
				"v_202917":               goqu.L("EXCLUDED.v_202917"),
				"m_202917":               goqu.L("EXCLUDED.m_202917"),
				"q_202917":               goqu.L("EXCLUDED.q_202917"),
				"v_202918":               goqu.L("EXCLUDED.v_202918"),
				"m_202918":               goqu.L("EXCLUDED.m_202918"),
				"q_202918":               goqu.L("EXCLUDED.q_202918"),
				"v_202919":               goqu.L("EXCLUDED.v_202919"),
				"m_202919":               goqu.L("EXCLUDED.m_202919"),
				"q_202919":               goqu.L("EXCLUDED.q_202919"),
				"v_202920":               goqu.L("EXCLUDED.v_202920"),
				"m_202920":               goqu.L("EXCLUDED.m_202920"),
				"q_202920":               goqu.L("EXCLUDED.q_202920"),
				"v_202921":               goqu.L("EXCLUDED.v_202921"),
				"m_202921":               goqu.L("EXCLUDED.m_202921"),
				"q_202921":               goqu.L("EXCLUDED.q_202921"),
				"v_202932":               goqu.L("EXCLUDED.v_202932"),
				"m_202932":               goqu.L("EXCLUDED.m_202932"),
				"q_202932":               goqu.L("EXCLUDED.q_202932"),
				"v_202935":               goqu.L("EXCLUDED.v_202935"),
				"m_202935":               goqu.L("EXCLUDED.m_202935"),
				"q_202935":               goqu.L("EXCLUDED.q_202935"),
				"v_202924":               goqu.L("EXCLUDED.v_202924"),
				"m_202924":               goqu.L("EXCLUDED.m_202924"),
				"q_202924":               goqu.L("EXCLUDED.q_202924"),
				"v_202925":               goqu.L("EXCLUDED.v_202925"),
				"m_202925":               goqu.L("EXCLUDED.m_202925"),
				"q_202925":               goqu.L("EXCLUDED.q_202925"),
				"v_203565":               goqu.L("EXCLUDED.v_203565"),
				"m_203565":               goqu.L("EXCLUDED.m_203565"),
				"q_203565":               goqu.L("EXCLUDED.q_203565"),
				"v_209190":               goqu.L("EXCLUDED.v_209190"),
				"m_209190":               goqu.L("EXCLUDED.m_209190"),
				"q_209190":               goqu.L("EXCLUDED.q_209190"),
				"v_203570":               goqu.L("EXCLUDED.v_203570"),
				"m_203570":               goqu.L("EXCLUDED.m_203570"),
				"q_203570":               goqu.L("EXCLUDED.q_203570"),
				"v_203551":               goqu.L("EXCLUDED.v_203551"),
				"m_203551":               goqu.L("EXCLUDED.m_203551"),
				"q_203551":               goqu.L("EXCLUDED.q_203551"),
				"v_202936":               goqu.L("EXCLUDED.v_202936"),
				"m_202936":               goqu.L("EXCLUDED.m_202936"),
				"q_202936":               goqu.L("EXCLUDED.q_202936"),
				"v_203569":               goqu.L("EXCLUDED.v_203569"),
				"m_203569":               goqu.L("EXCLUDED.m_203569"),
				"q_203569":               goqu.L("EXCLUDED.q_203569"),
				"v_203557":               goqu.L("EXCLUDED.v_203557"),
				"m_203557":               goqu.L("EXCLUDED.m_203557"),
				"q_203557":               goqu.L("EXCLUDED.q_203557"),
				"v_203568":               goqu.L("EXCLUDED.v_203568"),
				"m_203568":               goqu.L("EXCLUDED.m_203568"),
				"q_203568":               goqu.L("EXCLUDED.q_203568"),
				"v_203559":               goqu.L("EXCLUDED.v_203559"),
				"m_203559":               goqu.L("EXCLUDED.m_203559"),
				"q_203559":               goqu.L("EXCLUDED.q_203559"),
				"v_203577":               goqu.L("EXCLUDED.v_203577"),
				"m_203577":               goqu.L("EXCLUDED.m_203577"),
				"q_203577":               goqu.L("EXCLUDED.q_203577"),
				"v_211082":               goqu.L("EXCLUDED.v_211082"),
				"m_211082":               goqu.L("EXCLUDED.m_211082"),
				"q_211082":               goqu.L("EXCLUDED.q_211082"),
				"v_202931":               goqu.L("EXCLUDED.v_202931"),
				"m_202931":               goqu.L("EXCLUDED.m_202931"),
				"q_202931":               goqu.L("EXCLUDED.q_202931"),
			}),
		)
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to generate query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "save measurement")
	}

	return nil
}

func (r *repository) DeleteCreatedBefore(ctx context.Context, before *time.Time) error {
	beforeParam := before.Format("2006-01-02 15:04:05")
	query := goqu.Dialect("postgres").
		Delete("measurements").
		Where(goqu.L("real_date_time::date <=", beforeParam))
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to generate query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "delete measurements")
	}

	return nil
}

// TODO: getLastPostMeasurement ?? get can change it with params: take = 1, offset = 0
