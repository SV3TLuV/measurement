package model

import "time"

type CollectorInformationView struct {
	ListenedPostCount   uint64     `json:"listenedPostCount"`
	PostCount           uint64     `json:"postCount"`
	PollingInterval     uint64     `json:"pollingInterval"`
	LastPollingDateTime *time.Time `json:"lastPollingDateTime"`
	UntilNextPolling    *time.Time `json:"untilNextPolling"`
}

type CollectorStateView struct {
	Status          string     `json:"status"`
	PolledPostCount uint64     `json:"polledPostCount"`
	PostCount       uint64     `json:"postCount"`
	PollingPercent  uint64     `json:"pollingPercent"`
	ReceivedCount   uint64     `json:"receivedCount"`
	Started         *time.Time `json:"started"`
}
