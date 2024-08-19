package model

type User struct {
	ID            uint64
	Login         string
	Password      string
	IsBlocked     bool
	Role          Role
	PermissionIds []uint64
	ColumnIds     []uint64
	PostIds       []uint64
}
