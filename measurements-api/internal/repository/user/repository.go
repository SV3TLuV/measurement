package user

import (
	"context"
	"fmt"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"measurements-api/internal/model"
	def "measurements-api/internal/repository"
	"measurements-api/internal/repository/column/converter"
	"measurements-api/internal/repository/entities"
	converter5 "measurements-api/internal/repository/object/converter"
	converter3 "measurements-api/internal/repository/permission/converter"
	converter2 "measurements-api/internal/repository/role/converter"
	converter4 "measurements-api/internal/repository/user/converter"
	model2 "measurements-api/internal/repository/user/model"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	pool      *pgxpool.Pool
	getter    *trmpgx.CtxGetter
	trManager *manager.Manager
}

func NewRepository(
	pool *pgxpool.Pool,
	getter *trmpgx.CtxGetter,
	trManager *manager.Manager) *repository {
	return &repository{
		pool:      pool,
		getter:    getter,
		trManager: trManager,
	}
}

func (r *repository) Get(ctx context.Context,
	options *model2.GetUsersQueryParams) ([]*model.User, *uint64, error) {
	if options == nil {
		return nil, nil, errors.Wrap(model.Empty, "options")
	}

	dialect := goqu.Dialect("postgres")
	query := dialect.From("users")

	if len(options.RoleIds) > 0 {
		query = query.Where(goqu.Ex{
			"users.role_id": goqu.Op{"in": options.RoleIds},
		})
	}
	if options.Search != nil {
		search := fmt.Sprintf("%s%%", *options.Search)
		query = query.Where(goqu.Ex{
			"users.login": goqu.Op{"ilike": search},
		})
	}

	uc := dialect.From(goqu.T("user_columns")).
		Select(goqu.C("user_id"), goqu.L("array_agg(column_id)").As("columnIds")).
		GroupBy(goqu.C("user_id"))

	up := dialect.From(goqu.T("user_permissions")).
		Select(goqu.C("user_id"), goqu.L("array_agg(permission_id)").As("permissionIds")).
		GroupBy(goqu.C("user_id"))

	u := dialect.From("user_posts").
		Select(goqu.C("user_id"), goqu.L("array_agg(object_id)").As("postIds")).
		GroupBy(goqu.C("user_id"))

	selectQuery := query.Clone().(*goqu.SelectDataset).
		Select(
			"users.*",
			goqu.I("roles.role_id").As(goqu.C("role.role_id")),
			goqu.I("roles.title").As(goqu.C("role.title")),
			goqu.I("roles.name").As(goqu.C("role.name")),
			goqu.COALESCE(goqu.I("uc.columnIds"), goqu.L("'{}'")).As("columnIds"),
			goqu.COALESCE(goqu.I("up.permissionIds"), goqu.L("'{}'")).As("permissionIds"),
			goqu.COALESCE(goqu.I("u.postIds"), goqu.L("'{}'")).As("postIds")).
		Join(
			goqu.T("roles"),
			goqu.On(goqu.Ex{"users.role_id": goqu.I("roles.role_id")})).
		LeftJoin(
			up.As("up"),
			goqu.On(goqu.Ex{"up.user_id": goqu.I("users.user_id")})).
		LeftJoin(
			uc.As("uc"),
			goqu.On(goqu.Ex{"uc.user_id": goqu.I("users.user_id")})).
		LeftJoin(
			u.As("u"),
			goqu.On(goqu.Ex{"u.user_id": goqu.I("users.user_id")})).
		Order(goqu.I("users.user_id").Asc()).
		GroupBy(
			goqu.I("users.user_id"),
			goqu.C("role.role_id"),
			goqu.C("role.title"),
			goqu.C("role.name"),
			goqu.I("uc.columnIds"),
			goqu.I("up.permissionIds"),
			goqu.I("u.postIds")).
		Limit(options.PageSize).
		Offset((options.Page - 1) * options.PageSize)

	countQuery := query.Clone().(*goqu.SelectDataset).
		Select(goqu.L("COUNT(*)").As("total_count"))

	sql, args, err := selectQuery.ToSQL()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate select sql")
	}

	var users []*entities.UserWithRelated
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &users, sql, args...)
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

	models := make([]*model.User, 0, len(users))
	for i := 0; i < len(users); i++ {
		user := &model.User{
			ID:            users[i].ID,
			Login:         users[i].Login,
			Password:      users[i].PasswordHash,
			IsBlocked:     users[i].IsBlocked,
			Role:          *converter2.ToRoleFromRepo(&users[i].Role),
			PermissionIds: users[i].PermissionIds,
			ColumnIds:     users[i].ColumnIds,
			PostIds:       users[i].PostIds,
		}

		models = append(models, user)
	}

	return models, &total, nil
}

func (r *repository) GetByIds(ctx context.Context, ids []uint64) ([]*model.User, error) {
	if len(ids) == 0 {
		return nil, errors.Wrap(model.Empty, "ids")
	}

	query := goqu.Dialect("postgres").
		From("users").
		Select(
			"users.*",
			goqu.I("roles.role_id").As(goqu.C("role.role_id")),
			goqu.I("roles.title").As(goqu.C("role.title")),
			goqu.I("roles.name").As(goqu.C("role.name"))).
		Join(
			goqu.T("roles"),
			goqu.On(goqu.Ex{"users.role_id": goqu.I("roles.role_id")})).
		Where(goqu.Ex{"user_id": goqu.Op{"in": ids}}).
		Order(goqu.I("user_id").Asc())
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql")
	}

	var users []*entities.User
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &users, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter4.ToUsersFromRepo(users), nil
}

func (r *repository) GetById(ctx context.Context, id uint64) (*model.User, error) {
	query := goqu.Dialect("postgres").
		From("users").
		Select(
			"users.*",
			goqu.I("roles.role_id").As(goqu.C("role.role_id")),
			goqu.I("roles.title").As(goqu.C("role.title")),
			goqu.I("roles.name").As(goqu.C("role.name"))).
		Join(
			goqu.T("roles"),
			goqu.On(goqu.Ex{"users.role_id": goqu.I("roles.role_id")})).
		Where(goqu.Ex{"user_id": id})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql")
	}

	var user entities.User
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &user, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter4.ToUserFromRepo(&user), nil
}

func (r *repository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	query := goqu.Dialect("postgres").
		From("users").
		Select(
			"users.*",
			goqu.I("roles.role_id").As(goqu.C("role.role_id")),
			goqu.I("roles.title").As(goqu.C("role.title")),
			goqu.I("roles.name").As(goqu.C("role.name"))).
		Join(
			goqu.T("roles"),
			goqu.On(goqu.Ex{"users.role_id": goqu.I("roles.role_id")})).
		Where(goqu.Ex{"login": login})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql")
	}

	var user entities.User
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Get(ctx, tr, &user, sql, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, model.NotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter4.ToUserFromRepo(&user), nil
}

func (r *repository) Save(ctx context.Context, user *model.User) error {
	record := converter4.ToUserRecordFromService(user)
	query := goqu.Dialect("postgres").
		Insert("users").
		Rows(record).
		OnConflict(
			goqu.DoUpdate("user_id", &goqu.Record{
				"login":         goqu.L("EXCLUDED.login"),
				"password_hash": goqu.L("EXCLUDED.password_hash"),
				"role_id":       goqu.L("EXCLUDED.role_id"),
				"is_blocked":    goqu.L("EXCLUDED.is_blocked"),
			}))
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "save user")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id uint64) error {
	query := goqu.Dialect("postgres").
		Delete("users").
		Where(goqu.Ex{"user_id": id})
	sql, args, err := query.ToSQL()
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}

	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	_, err = tr.Exec(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "delete user")
	}

	return nil
}

func (r *repository) GetUserPermissions(ctx context.Context, userID uint64) ([]*model.Permission, error) {
	query := goqu.Dialect("postgres").
		From("permissions").
		Select("permissions.*").
		Join(
			goqu.T("user_permissions"),
			goqu.On(goqu.Ex{"permissions.permission_id": goqu.I("user_permissions.permission_id")})).
		Where(goqu.Ex{"user_id": userID})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var permissions []*entities.Permission
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &permissions, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter3.ToPermissionsFromRepo(permissions), nil
}

func (r *repository) UpdateUserPermissions(ctx context.Context, userID uint64, permissionIDs []uint64) error {
	err := r.trManager.Do(ctx, func(ctx context.Context) error {
		deleteQuery := goqu.Dialect("postgres").
			Delete("user_permissions").
			Where(goqu.Ex{"user_id": userID})
		if len(permissionIDs) > 0 {
			deleteQuery = deleteQuery.
				Where(goqu.Ex{"permission_id": goqu.Op{"notIn": permissionIDs}})
		}

		sql, args, err := deleteQuery.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		tr := r.getter.DefaultTrOrDB(ctx, r.pool)
		_, err = tr.Exec(ctx, sql, args...)
		if err != nil {
			return errors.Wrap(err, "delete permissions")
		}
		if len(permissionIDs) == 0 {
			return nil
		}

		records := make([]*goqu.Record, 0, len(permissionIDs))
		for i := 0; i < len(permissionIDs); i++ {
			records = append(records, &goqu.Record{
				"user_id":       userID,
				"permission_id": permissionIDs[i],
			})
		}

		insertQuery := goqu.Dialect("postgres").
			Insert("user_permissions").
			Rows(records)
		sql, args, err = insertQuery.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		_, err = tr.Exec(ctx, sql, args...)
		if err != nil {
			return errors.Wrap(err, "insert permissions")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}

	return nil
}

func (r *repository) GetUserColumns(ctx context.Context, userID uint64) ([]*model.Column, error) {
	query := goqu.Dialect("postgres").
		From("columns").
		Select("columns.*").
		Join(
			goqu.T("user_columns"),
			goqu.On(goqu.Ex{"columns.column_id": goqu.I("user_columns.column_id")})).
		Where(goqu.Ex{"user_id": userID})
	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}

	var columns []*entities.Column
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err = pgxscan.Select(ctx, tr, &columns, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	return converter.ToColumnsFromRepo(columns), nil
}

func (r *repository) UpdateUserColumns(ctx context.Context, userID uint64, columnIDs []uint64) error {
	err := r.trManager.Do(ctx, func(ctx context.Context) error {
		deleteQuery := goqu.Dialect("postgres").
			Delete("user_columns").
			Where(goqu.Ex{"user_id": userID})
		if len(columnIDs) > 0 {
			deleteQuery = deleteQuery.
				Where(goqu.Ex{"column_id": goqu.Op{"notIn": columnIDs}})
		}

		sql, args, err := deleteQuery.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		tr := r.getter.DefaultTrOrDB(ctx, r.pool)
		_, err = tr.Exec(ctx, sql, args...)
		if err != nil {
			return errors.Wrap(err, "delete columns")
		}
		if len(columnIDs) == 0 {
			return nil
		}

		records := make([]*goqu.Record, 0, len(columnIDs))
		for i := 0; i < len(columnIDs); i++ {
			records = append(records, &goqu.Record{
				"user_id":   userID,
				"column_id": columnIDs[i],
			})
		}

		insertQuery := goqu.Dialect("postgres").
			Insert("user_columns").
			Rows(records)
		sql, args, err = insertQuery.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		_, err = tr.Exec(ctx, sql, args...)
		if err != nil {
			return errors.Wrap(err, "insert columns")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}

	return nil
}

func (r *repository) GetUserObjects(ctx context.Context, userID uint64) ([]*model.Object, error) {
	const query = `
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
		), object_hierarchy AS (
			SELECT
				o.*,
				up.user_id,
				ot.object_type_id AS "type.object_type_id",
				ot.title AS "type.title",
				pi.object_id AS "post_info.object_id",
				pi.last_polling_date_time AS "post_info.last_polling_date_time",
				pi.is_listened AS "post_info.is_listened"
			FROM objects o
			JOIN user_posts up ON up.object_id = o.object_id
			JOIN object_types ot ON ot.object_type_id = o.type_id
			LEFT JOIN post_infos pi ON pi.object_id = o.object_id
			WHERE o.type_id = 3 AND up.user_id = $1 AND pi.is_listened
		
			UNION ALL
		
			SELECT
				o.*,
				oh.user_id,
				ot.object_type_id AS "type.object_type_id",
				ot.title AS "type.title",
				pi.object_id AS "post_info.object_id",
				pi.last_polling_date_time AS "post_info.last_polling_date_time",
				pi.is_listened AS "post_info.is_listened"
			FROM object_hierarchy AS oh
			JOIN objects o ON o.object_id = oh.parent_id
			JOIN object_types ot ON ot.object_type_id = o.type_id
			LEFT JOIN post_infos pi ON pi.object_id = o.object_id
		)
		
		SELECT DISTINCT
			oh.object_id,
			oh.parent_id,
			oh.type_id,
			oh.title,
			oh.address,
			oh.lat,
			oh.lon,
			CASE WHEN oh.type_id = 1 THEN NULL ELSE ot.laboratory_title END AS laboratory_title,
			CASE WHEN oh.type_id = 2 THEN NULL ELSE ot.city_title END AS city_title,
			oh."type.object_type_id",
			oh."type.title",
			oh."post_info.object_id",
			oh."post_info.last_polling_date_time",
			oh."post_info.is_listened"
		FROM object_hierarchy AS oh
		JOIN object_titles ot ON ot.object_id = oh.object_id
		ORDER BY oh.object_id
	`

	var objects []*entities.Object
	tr := r.getter.DefaultTrOrDB(ctx, r.pool)
	err := pgxscan.Select(ctx, tr, &objects, query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "select")
	}

	objectMap := make(map[uint64]*entities.Object)
	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		objectMap[obj.ID.Int.Uint64()] = obj
	}

	for i := 0; i < len(objects); i++ {
		obj := objects[i]
		if obj.ParentID.Valid {
			parent := objectMap[obj.ParentID.Int.Uint64()]
			if parent != nil {
				parent.Children = append(parent.Children, obj)
			}
		}
	}

	rootObjects := make([]*entities.Object, 0)
	for _, obj := range objectMap {
		if !obj.ParentID.Valid {
			rootObjects = append(rootObjects, obj)
		}
	}

	return converter5.ToObjectsFromRepo(rootObjects), nil
}

func (r *repository) UpdateUserPosts(ctx context.Context, userID uint64, postIDs []uint64) error {
	err := r.trManager.Do(ctx, func(ctx context.Context) error {
		deleteQuery := goqu.Dialect("postgres").
			Delete("user_posts").
			Where(goqu.Ex{"user_id": userID})
		if len(postIDs) > 0 {
			deleteQuery = deleteQuery.
				Where(goqu.Ex{"object_id": goqu.Op{"notIn": postIDs}})
		}

		sql, args, err := deleteQuery.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		tr := r.getter.DefaultTrOrDB(ctx, r.pool)
		_, err = tr.Exec(ctx, sql, args...)
		if err != nil {
			return errors.Wrap(err, "delete posts")
		}
		if len(postIDs) == 0 {
			return nil
		}

		records := make([]*goqu.Record, 0, len(postIDs))
		for i := 0; i < len(postIDs); i++ {
			records = append(records, &goqu.Record{
				"user_id":   userID,
				"object_id": postIDs[i],
			})
		}

		insertQuery := goqu.Dialect("postgres").
			Insert("user_posts").
			Rows(records)
		sql, args, err = insertQuery.ToSQL()
		if err != nil {
			return errors.Wrap(err, "failed to build query")
		}

		_, err = tr.Exec(ctx, sql, args...)
		if err != nil {
			return errors.Wrap(err, "insert posts")
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction failed")
	}

	return nil
}
