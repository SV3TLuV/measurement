package model

import (
	"time"
)

type Object struct {
	ID         uint64
	Title      string
	Address    *string
	Lat        *float64
	Lon        *float64
	Type       ObjectType
	PostInfo   *PostInfo
	ParentID   *uint64
	Laboratory *string
	City       *string
	Status     *string
	Children   []*Object
}

type PostInfo struct {
	ObjectID            uint64
	LastPollingDateTime *time.Time
	IsListened          bool
}
