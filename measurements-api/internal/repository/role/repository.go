package role

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
	"measurements-api/internal/repository/role/converter"
)

var _ def.RoleRepository = (*repository)(nil)

type repository struct {
	pool   *pgxpool.Pool
	getter *trmpgx.CtxGetter
}

func NewRepository(db *pgxpool.Pool, c *trmpgx.CtxGetter) *repository {
	return &repository{
		pool:   db,
		getter: c,
	}
}

func (r *repository) Get(ctx context.Context) ([]*model.Role, error) {
	query := goqu.Dialect("postgres").From("roles")
	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to generate sql")
	}

	var roles []*entities.Role
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &roles, sql)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToRolesFromRepo(roles), nil
}
