package model

import "time"

type PollingStatisticView struct {
	ID            uint64    `json:"id"`
	DateTime      time.Time `json:"dateTime"`
	Duration      string    `json:"duration"`
	PostCount     uint64    `json:"postCount"`
	ReceivedCount uint64    `json:"receivedCount"`
}

type FetchPollingStatisticQuery struct {
	Page     uint `query:"page" validate:"gte=1"`
	PageSize uint `query:"pageSize" validate:"gte=1"`
}
