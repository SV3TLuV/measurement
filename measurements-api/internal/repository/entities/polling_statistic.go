package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type PollingStatistic struct {
	ID            uint64          `db:"polling_statistic_id"`
	DateTime      time.Time       `db:"datetime"`
	Duration      pgtype.Interval `db:"duration"`
	PostCount     uint64          `db:"post_count"`
	ReceivedCount uint64          `db:"received_count"`
}
