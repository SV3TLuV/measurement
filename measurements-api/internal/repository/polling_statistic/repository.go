package polling_statistic

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	def "measurements-api/internal/repository"
	"measurements-api/internal/repository/entities"
	"measurements-api/internal/repository/polling_statistic/converter"
	model2 "measurements-api/internal/repository/polling_statistic/model"
	"slices"
)

var _ def.PollingStatisticRepository = (*repository)(nil)

type repository struct {
	pool   *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewRepository(
	pool *pgxpool.Pool,
	getter *trmpgx.CtxGetter) *repository {
	return &repository{
		pool:   pool,
		getter: getter,
	}
}

func (r *repository) Get(ctx context.Context,
	options *model2.GetPollingStatisticParams) ([]*model.PollingStatistic, error) {
	query := goqu.Dialect("postgres").
		From("polling_statistics").
		Order(goqu.I("polling_statistic_id").Desc())

	if options != nil {
		query = query.Limit(options.PageSize).
			Offset((options.Page - 1) * options.PageSize)
	}

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate select sql")
	}

	var statistics []*entities.PollingStatistic
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &statistics, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	slices.SortFunc(statistics, func(a, b *entities.PollingStatistic) int {
		return int(a.ID - b.ID)
	})

	return converter.ToPollingStatisticsFromRepo(statistics), err
}

func (r *repository) SaveOne(ctx context.Context, statistic *model.PollingStatistic) error {
	record := converter.ToPollingStatisticRecordFromService(statistic)
	query := goqu.Dialect("postgres").
		Insert("polling_statistics").
		Rows(record)
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to generate query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "save statistic")
	}

	return nil
}
