package model

import (
	"measurements-api/internal/server/http/controller/role/model"
)

type UserView struct {
	ID            uint64         `json:"id"`
	Login         string         `json:"login"`
	IsBlocked     bool           `json:"isBlocked"`
	Role          model.RoleView `json:"role"`
	PermissionIds []uint64       `json:"permissions,omitempty"`
	ColumnIds     []uint64       `json:"columns,omitempty"`
	PostIds       []uint64       `json:"posts,omitempty"`
}

type UserRequestWithID struct {
	ID uint64 `param:"id" validate:"required,gte=1"`
}

type FetchUserListRequest struct {
	Search   *string  `query:"search"`
	Page     uint     `query:"page" validate:"gte=1"`
	PageSize uint     `query:"pageSize" validate:"gte=1"`
	RoleIds  []uint64 `query:"roleIds"`
}

type CreateUserRequest struct {
	Login         string   `form:"login" validate:"required,gte=8"`
	Password      string   `form:"password" validate:"required,gte=12"`
	RoleID        uint64   `form:"roleId" validate:"required,gte=1"`
	PermissionIDs []uint64 `form:"permissionIds"`
	ColumnIDs     []uint64 `form:"columnIds"`
	PostIDs       []uint64 `form:"postIds"`
}

type UpdateUserRequest struct {
	UserID        uint64   `form:"userId" validate:"required,gte=1"`
	Login         *string  `form:"login" validate:"gte=8"`
	RoleID        *uint64  `form:"roleId" validate:"gte=1"`
	PermissionIDs []uint64 `form:"permissionIds"`
	ColumnIDs     []uint64 `form:"columnIds"`
	PostIDs       []uint64 `form:"postIds"`
}

type ChangePasswordRequest struct {
	UserID   uint64 `form:"userId" validate:"required,gte=1"`
	Password string `form:"password" validate:"required,gte=12"`
}
