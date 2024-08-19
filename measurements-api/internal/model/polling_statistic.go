package model

import "time"

type PollingStatistic struct {
	ID            uint64
	DateTime      time.Time
	Duration      time.Duration
	PostCount     uint64
	ReceivedCount uint64
}
