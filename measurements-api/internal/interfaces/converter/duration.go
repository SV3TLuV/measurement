package converter

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func ToDurationFromInterval(interval pgtype.Interval) time.Duration {
	duration := time.Duration(interval.Days) * 24 * time.Hour
	duration += time.Duration(interval.Microseconds) * time.Microsecond
	return duration
}

func ToIntervalFromDuration(duration time.Duration) pgtype.Interval {
	totalMicroseconds := duration.Microseconds()
	days := totalMicroseconds / (24 * 60 * 60 * 1e6)
	remainingMicroseconds := totalMicroseconds % (24 * 60 * 60 * 1e6)
	return pgtype.Interval{
		Days:         int32(days),
		Microseconds: remainingMicroseconds,
	}
}
