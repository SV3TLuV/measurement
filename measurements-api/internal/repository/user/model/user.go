package model

type GetUsersQueryParams struct {
	Search   *string
	Page     uint
	PageSize uint
	RoleIds  []uint64
}
