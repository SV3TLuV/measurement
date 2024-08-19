package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type PostInfo struct {
	ObjectID            pgtype.Numeric `db:"object_id"`
	LastPollingDateTime *time.Time     `db:"last_polling_date_time"`
	IsListened          pgtype.Bool    `db:"is_listened"`
}
