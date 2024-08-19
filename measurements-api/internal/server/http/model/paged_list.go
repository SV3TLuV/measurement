package model

type PagedListView[T any] struct {
	Page     uint   `json:"page"`
	PageSize uint   `json:"pageSize"`
	Total    uint64 `json:"total"`
	Items    []*T   `json:"items"`
}
