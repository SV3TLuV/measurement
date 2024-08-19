package model

import "time"

type ObjectView struct {
	ID                  uint64          `json:"id"`
	Title               string          `json:"title"`
	Address             *string         `json:"address"`
	Lat                 *float64        `json:"lat"`
	Lon                 *float64        `json:"lon"`
	Type                *ObjectTypeView `json:"type"`
	ParentID            *uint64         `json:"parentId"`
	LastPollingDateTime *time.Time      `json:"lastPollingDateTime,omitempty"`
	IsListened          *bool           `json:"isListened,omitempty"`
	Laboratory          *string         `json:"laboratory,omitempty"`
	City                *string         `json:"city,omitempty"`
	Status              *string         `json:"status,omitempty"`
	Children            []*ObjectView   `json:"children"`
}

type ObjectTypeView struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type FetchObjectListQuery struct {
	TypeId    *uint64  `query:"typeId"`
	Search    *string  `query:"search"`
	ParentIds []uint64 `query:"parentIds"`
}

type ObjectRequestWithID struct {
	ID uint64 `param:"id" validate:"required,gte=1"`
}
