package entities

import "github.com/jackc/pgx/v5/pgtype"

type Permission struct {
	ID    pgtype.Numeric `db:"permission_id"`
	Name  pgtype.Text    `db:"name"`
	Title pgtype.Text    `db:"title"`
}
