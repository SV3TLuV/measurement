package session

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	def "measurements-api/internal/repository"
	"measurements-api/internal/repository/entities"
	"measurements-api/internal/repository/session/converter"
)

var _ def.SessionRepository = (*repository)(nil)

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

func (r *repository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]*model.Session, error) {
	if len(ids) == 0 {
		return nil, errors.Wrap(model.Empty, "ids")
	}

	query := goqu.Dialect("postgres").
		From("sessions").
		Where(goqu.Ex{
			"session_id": goqu.Op{"in": ids},
		})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var sessions []*entities.Session
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &sessions, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToSessionsFromRepo(sessions), nil
}

func (r *repository) GetById(ctx context.Context, id uuid.UUID) (*model.Session, error) {
	query := goqu.Dialect("postgres").
		From("sessions").
		Where(goqu.Ex{"session_id": id.String()})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to generate sql")
	}

	var session entities.Session
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &session, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToSessionFromRepo(&session), nil
}

func (r *repository) Save(ctx context.Context, session *model.Session) error {
	record := converter.ToSessionRecordFromService(session)
	query := goqu.Dialect("postgres").
		Insert("sessions").
		Rows(record).
		OnConflict(
			goqu.DoUpdate("session_id", &goqu.Record{
				"updated":       goqu.L("EXCLUDED.updated"),
				"refresh_token": goqu.L("EXCLUDED.refresh_token"),
			}))
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "save")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return errors.Wrap(model.Empty, "ids")
	}

	query := goqu.Dialect("postgres").
		Delete("sessions").
		Where(goqu.Ex{
			"session_id": goqu.Op{"in": ids},
		})
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "delete")
	}

	return nil
}

func (r *repository) DeleteOne(ctx context.Context, id uuid.UUID) error {
	query := goqu.Dialect("postgres").
		Delete("sessions").
		Where(goqu.Ex{
			"session_id": id,
		})
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "delete")
	}

	return nil
}
