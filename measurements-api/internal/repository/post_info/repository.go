package post_info

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
	"measurements-api/internal/repository/post_info/converter"
)

var _ def.PostInfoRepository = (*repository)(nil)

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

func (r *repository) GetById(ctx context.Context, id uint64) (*model.PostInfo, error) {
	query := goqu.Dialect("postgres").
		From("post_infos").
		Where(goqu.Ex{"object_id": id})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var postInfo entities.PostInfo
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &postInfo, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToPostInfoFromRepo(&postInfo), err
}

func (r *repository) Save(ctx context.Context, postInfos []*model.PostInfo) error {
	if len(postInfos) == 0 {
		return nil
	}

	records := converter.ToPostInfoRecordsFromService(postInfos)
	query := goqu.Dialect("postgres").
		Insert("post_infos").
		Rows(records).
		OnConflict(
			goqu.DoUpdate("object_id", &goqu.Record{
				"last_polling_date_time": goqu.L("EXCLUDED.last_polling_date_time"),
				"is_listened":            goqu.L("EXCLUDED.is_listened"),
			}),
		)
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "save post_infos")
	}

	return nil
}

func (r *repository) SaveOne(ctx context.Context, postInfo *model.PostInfo) error {
	record := converter.ToPostInfoRecordFromService(postInfo)
	query := goqu.Dialect("postgres").
		Insert("post_infos").
		Rows(record).
		OnConflict(
			goqu.DoUpdate("object_id", &goqu.Record{
				"last_polling_date_time": goqu.L("EXCLUDED.last_polling_date_time"),
				"is_listened":            goqu.L("EXCLUDED.is_listened"),
			}),
		)
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "save post_info")
	}

	return nil
}
