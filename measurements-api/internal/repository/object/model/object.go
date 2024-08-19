package model

type GetObjectsQueryParams struct {
	TypeID    *uint64
	Search    *string
	ParentIds []uint64
}

type GetObjectCountParams struct {
	TypeID     *uint64
	IsListened *bool
}
