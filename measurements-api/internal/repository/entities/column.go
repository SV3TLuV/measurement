package entities

import "github.com/jackc/pgx/v5/pgtype"

type Column struct {
	ID         pgtype.Numeric `db:"column_id"`
	Title      pgtype.Text    `db:"title"`
	ShortTitle pgtype.Text    `db:"short_title"`
	Formula    pgtype.Text    `db:"formula"`
	ObjField   pgtype.Text    `db:"obj_field"`
	Code       pgtype.Text    `db:"code"`
}
