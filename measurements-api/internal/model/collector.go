package model

import "time"

const (
	Pending    CollectorStatus = "pending"
	Collecting CollectorStatus = "collecting"
)

type CollectorStatus string

type CollectorState struct {
	Status          CollectorStatus `json:"status"`
	PolledPostCount uint64          `json:"polled_post_count"`
	PostCount       uint64          `json:"post_count"`
	PollingPercent  uint64          `json:"polling_percent"`
	ReceivedCount   uint64          `json:"received_count"`
	Started         *time.Time      `json:"started"`
	Ended           *time.Time      `json:"ended"`
}

type CollectorInformation struct {
	ListenedPostCount   uint64
	PostCount           uint64
	PollingInterval     uint64
	LastPollingDateTime *time.Time
	UntilNextPolling    *time.Time
}
