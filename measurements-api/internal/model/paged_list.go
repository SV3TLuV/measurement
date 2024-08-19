package model

type PagedList[T any] struct {
	Page     uint
	PageSize uint
	Total    uint64
	Items    []*T
}
