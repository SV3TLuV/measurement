package entities

import "github.com/jackc/pgx/v5/pgtype"

type Object struct {
	ID         pgtype.Numeric `db:"object_id"`
	Title      pgtype.Text    `db:"title"`
	Address    pgtype.Text    `db:"address"`
	Lat        pgtype.Float8  `db:"lat"`
	Lon        pgtype.Float8  `db:"lon"`
	TypeID     pgtype.Numeric `db:"type_id"`
	ParentID   pgtype.Numeric `db:"parent_id"`
	Type       ObjectType     `db:"type"`
	PostInfo   *PostInfo      `db:"post_info"`
	Laboratory pgtype.Text    `db:"laboratory_title"`
	City       pgtype.Text    `db:"city_title"`
	Children   []*Object
}

type ObjectWithOperation struct {
	Object
	Operation string `db:"operation"`
}

type ObjectType struct {
	ID    uint64 `db:"object_type_id"`
	Title string `db:"title"`
}
