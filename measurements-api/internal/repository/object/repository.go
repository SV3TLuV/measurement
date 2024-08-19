package object

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
	"measurements-api/internal/repository/object/converter"
	model2 "measurements-api/internal/repository/object/model"
)

var _ def.ObjectRepository = (*repository)(nil)

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
	options *model2.GetObjectsQueryParams) ([]*model.Object, error) {
	const objectTitlesQuery = `
		WITH RECURSIVE object_infos AS (
			SELECT
				o.object_id,
				o.object_id AS laboratory_id,
				NULL::int AS city_id,
				o.title AS laboratory_title,
				NULL AS city_title
			FROM objects o
			WHERE o.type_id = 1 AND o.parent_id IS NULL
		
			UNION ALL
		
			SELECT
				o.object_id,
				oi.laboratory_id,
				CASE WHEN o.type_id = 2 THEN o.object_id ELSE oi.city_id END AS city_id,
				oi.laboratory_title,
				CASE WHEN o.type_id = 2 THEN o.title ELSE oi.city_title END AS city_title
			FROM object_infos oi
			JOIN objects o ON o.parent_id = oi.object_id
		), sorted_object_infos AS (
			SELECT *
			FROM object_infos
			ORDER BY laboratory_id, city_id, object_id
		)
	`

	query := goqu.Dialect("postgres").
		From("objects").
		Select(
			"objects.*",
			goqu.I("sorted_object_infos.laboratory_title"),
			goqu.I("sorted_object_infos.city_title"),
			goqu.I("object_types.object_type_id").As(goqu.C("type.object_type_id")),
			goqu.I("object_types.title").As(goqu.C("type.title")),
			goqu.I("post_infos.object_id").As(goqu.C("post_info.object_id")),
			goqu.I("post_infos.last_polling_date_time").As(goqu.C("post_info.last_polling_date_time")),
			goqu.I("post_infos.is_listened").As(goqu.C("post_info.is_listened")),
		).
		Join(
			goqu.T("object_types"),
			goqu.On(goqu.Ex{"objects.type_id": goqu.I("object_types.object_type_id")})).
		Join(
			goqu.T("sorted_object_infos"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("sorted_object_infos.object_id")})).
		LeftJoin(
			goqu.T("post_infos"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("post_infos.object_id")}))

	if options != nil {
		if options.TypeID != nil {
			query = query.Where(goqu.Ex{"objects.type_id": *options.TypeID})
		}
		if options.Search != nil {
			search := fmt.Sprintf("%s%%", *options.Search)
			query = query.Where(
				goqu.Or(
					goqu.Ex{"objects.title": goqu.Op{"ilike": search}},
					goqu.Ex{"objects.address": goqu.Op{"ilike": search}},
					goqu.Ex{"sorted_object_infos.city_title": goqu.Op{"ilike": search}},
					goqu.Ex{"sorted_object_infos.laboratory_title": goqu.Op{"ilike": search}},
				))
		}
		if len(options.ParentIds) > 0 {
			query = query.Where(goqu.Or(
				goqu.Ex{"sorted_object_infos.laboratory_id": goqu.Op{"in": options.ParentIds}},
				goqu.Ex{"sorted_object_infos.city_id": goqu.Op{"in": options.ParentIds}}))
		}
	}

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql")
	}

	sql = fmt.Sprintf("%s %s", objectTitlesQuery, sql)

	var objects []*entities.Object
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &objects, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToObjectsFromRepo(objects), nil
}

func (r *repository) GetByIds(ctx context.Context, ids []uint64) ([]*model.Object, error) {
	if len(ids) == 0 {
		return nil, errors.Wrap(model.Empty, "ids")
	}

	const objectTitlesQuery = `
		WITH RECURSIVE object_titles AS (
			SELECT
				o.object_id,
				o.title AS laboratory_title,
				NULL AS city_title
			FROM objects o
			WHERE o.type_id = 1 AND o.parent_id IS NULL
		
			UNION ALL
		
			SELECT
				o.object_id,
				ot.laboratory_title,
				CASE WHEN o.type_id = 2 THEN o.title ELSE ot.city_title END AS city_title
			FROM object_titles 
			JOIN objects o ON o.parent_id = ot.object_id
		)
	`

	query := goqu.Dialect("postgres").
		From("objects").
		Select(
			"objects.*",
			goqu.I("object_titles.laboratory_title"),
			goqu.I("object_titles.city_title"),
			goqu.I("object_types.object_type_id").As(goqu.C("type.object_type_id")),
			goqu.I("object_types.title").As(goqu.C("type.title")),
			goqu.I("post_infos.object_id").As(goqu.C("post_info.object_id")),
			goqu.I("post_infos.last_polling_date_time").As(goqu.C("post_info.last_polling_date_time")),
			goqu.I("post_infos.is_listened").As(goqu.C("post_info.is_listened")),
		).
		Join(
			goqu.T("object_types"),
			goqu.On(goqu.Ex{"objects.type_id": goqu.I("object_types.object_type_id")})).
		Join(
			goqu.T("object_titles"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("object_titles.object_id")})).
		LeftJoin(
			goqu.T("post_infos"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("post_infos.object_id")})).
		Where(goqu.Ex{"object_id": goqu.Op{"in": ids}})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql")
	}

	sql = fmt.Sprintf("%s %s", objectTitlesQuery, sql)

	var objects []*entities.Object
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &objects, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToObjectsFromRepo(objects), nil
}

func (r *repository) GetById(ctx context.Context, id uint64) (*model.Object, error) {
	const objectTitlesQuery = `
		WITH RECURSIVE object_titles AS (
			SELECT
				o.object_id,
				o.title AS laboratory_title,
				NULL AS city_title
			FROM objects o
			WHERE o.type_id = 1 AND o.parent_id IS NULL
		
			UNION ALL
		
			SELECT
				o.object_id,
				ot.laboratory_title,
				CASE WHEN o.type_id = 2 THEN o.title ELSE ot.city_title END AS city_title
			FROM object_titles ot
			JOIN objects o ON o.parent_id = ot.object_id
		)
	`

	query := goqu.Dialect("postgres").
		From("objects").
		Select(
			"objects.*",
			goqu.I("object_titles.laboratory_title"),
			goqu.I("object_titles.city_title"),
			goqu.I("object_types.object_type_id").As(goqu.C("type.object_type_id")),
			goqu.I("object_types.title").As(goqu.C("type.title")),
			goqu.I("post_infos.object_id").As(goqu.C("post_info.object_id")),
			goqu.I("post_infos.last_polling_date_time").As(goqu.C("post_info.last_polling_date_time")),
			goqu.I("post_infos.is_listened").As(goqu.C("post_info.is_listened")),
		).
		Join(
			goqu.T("object_types"),
			goqu.On(goqu.Ex{"objects.type_id": goqu.I("object_types.object_type_id")})).
		Join(
			goqu.T("object_titles"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("object_titles.object_id")})).
		LeftJoin(
			goqu.T("post_infos"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("post_infos.object_id")})).
		Where(goqu.Ex{"objects.object_id": id})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql")
	}

	sql = fmt.Sprintf("%s %s", objectTitlesQuery, sql)

	var object entities.Object
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &object, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToObjectFromRepo(&object), nil
}

func (r *repository) GetUserPostById(ctx context.Context, userID, postID uint64) (*model.Object, error) {
	query := goqu.Dialect("postgres").
		From("objects").
		Select(
			"objects.*",
			goqu.I("object_types.object_type_id").As(goqu.C("type.object_type_id")),
			goqu.I("object_types.title").As(goqu.C("type.title")),
			goqu.I("post_infos.object_id").As(goqu.C("post_info.object_id")),
			goqu.I("post_infos.last_polling_date_time").As(goqu.C("post_info.last_polling_date_time")),
			goqu.I("post_infos.is_listened").As(goqu.C("post_info.is_listened"))).
		Join(
			goqu.T("object_types"),
			goqu.On(goqu.Ex{"objects.type_id": goqu.I("object_types.object_type_id")})).
		Join(
			goqu.T("user_posts"),
			goqu.On(goqu.Ex{"user_posts.object_id": goqu.I("objects.object_id")})).
		LeftJoin(
			goqu.T("post_infos"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("post_infos.object_id")})).
		Where(
			goqu.Ex{"objects.object_id": postID},
			goqu.Ex{"user_posts.user_id": userID})

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql")
	}

	var object entities.Object
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &object, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToObjectFromRepo(&object), nil
}

func (r *repository) GetCount(ctx context.Context, options *model2.GetObjectCountParams) (uint64, error) {
	query := goqu.Dialect("postgres").
		From("objects").
		Select(goqu.L("COUNT(*)")).
		LeftJoin(
			goqu.T("post_infos"),
			goqu.On(goqu.Ex{"objects.object_id": goqu.I("post_infos.object_id")}))

	if options != nil {
		if options.TypeID != nil {
			query = query.Where(goqu.Ex{"objects.type_id": *options.TypeID})
		}
		if options.IsListened != nil {
			query = query.Where(goqu.Ex{"post_infos.is_listened": *options.IsListened})
		}
	}

	sql, args, err := query.ToSQL()
	if err != nil {
		return 0, errors.Wrap(err, "failed to generate sql")
	}

	var count uint64
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &count, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "select")
	}

	return count, nil
}

func (r *repository) Save(ctx context.Context, objects []*model.Object) ([]*model.Object, error) {
	if len(objects) == 0 {
		return nil, nil
	}

	records := make([]*goqu.Record, 0, len(objects))
	for i := 0; i < len(objects); i++ {
		object := objects[i]
		records = append(records, &goqu.Record{
			"object_id": object.ID,
			"parent_od": object.ParentID,
			"type_id":   object.Type.ID,
			"title":     object.Title,
			"address":   object.Address,
			"lat":       object.Lat,
			"lon":       object.Lon,
		})
	}

	query := goqu.Dialect("postgres").
		Insert("objects").
		Rows(records).
		OnConflict(goqu.DoUpdate("object_id", &goqu.Record{
			"parent_id": goqu.L("EXCLUDED.parent_id"),
			"type_id":   goqu.L("EXCLUDED.type_id"),
			"title":     goqu.L("EXCLUDED.title"),
			"address":   goqu.L("EXCLUDED.address"),
			"lat":       goqu.L("EXCLUDED.lat"),
			"lon":       goqu.L("EXCLUDED.lon"),
		}).Where(
			goqu.Or(
				goqu.C("objects.parent_id").Neq(goqu.L("EXCLUDED.parent_id")),
				goqu.C("objects.type_id").Neq(goqu.L("EXCLUDED.type_id")),
				goqu.C("objects.title").Neq(goqu.L("EXCLUDED.title")),
				goqu.C("objects.address").Neq(goqu.L("EXCLUDED.address")),
				goqu.C("objects.lat").Neq(goqu.L("EXCLUDED.lat")),
				goqu.C("objects.lon").Neq(goqu.L("EXCLUDED.lon")),
			),
		)).
		Returning(
			"object_id", "parent_id", "type_id", "title", "address", "lat", "lon",
			goqu.L(`CASE
                    		WHEN EXCLUDED.object_id IS NOT NULL THEN 'updated'
                    		ELSE 'inserted'
                		END`).As("operation"))
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var objs []*entities.ObjectWithOperation
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &objs, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "save objects")
	}

	return converter.ToObjectsWithOperationFromRepo(objs), nil
}
