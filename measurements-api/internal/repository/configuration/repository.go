package configuration

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	def "measurements-api/internal/repository"
	"measurements-api/internal/repository/configuration/converter"
	"measurements-api/internal/repository/entities"
)

var _ def.ConfigurationRepository = (*repository)(nil)

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

func (r *repository) Get(ctx context.Context) (*model.Configuration, error) {
	const configurationID = 1
	query := goqu.Dialect("postgres").
		From("configurations").
		Where(goqu.Ex{"configuration_id": configurationID})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to generate sql")
	}

	var configuration entities.Configuration
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &configuration, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToConfigurationFromRepo(&configuration), err
}

func (r *repository) Save(ctx context.Context, configuration *model.Configuration) error {
	record := converter.ToConfigurationRecordFromService(configuration)
	query := goqu.Dialect("postgres").
		Insert("configurations").
		Rows(record).
		OnConflict(
			goqu.DoUpdate("configuration_id", &goqu.Record{
				"asoiza_login":        goqu.L("EXCLUDED.asoiza_login"),
				"asoiza_password":     goqu.L("EXCLUDED.asoiza_password"),
				"collecting_interval": goqu.L("EXCLUDED.collecting_interval"),
				"deleting_interval":   goqu.L("EXCLUDED.deleting_interval"),
				"deleting_threshold":  goqu.L("EXCLUDED.deleting_threshold"),
				"disabling_interval":  goqu.L("EXCLUDED.disabling_interval"),
				"disabling_threshold": goqu.L("EXCLUDED.disabling_threshold"),
			}),
		)
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "save configuration")
	}

	return nil
}
