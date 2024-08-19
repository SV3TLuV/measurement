package entities

type User struct {
	ID           uint64       `db:"user_id"`
	Login        string       `db:"login"`
	PasswordHash string       `db:"password_hash"`
	IsBlocked    bool         `db:"is_blocked"`
	RoleID       uint64       `db:"role_id"`
	Role         Role         `db:"role"`
	Permissions  []Permission `db:"permissions"`
	Columns      []Column     `db:"columns"`
	Posts        []Object     `db:"objects"`
}

type UserWithRelated struct {
	ID            uint64   `db:"user_id"`
	Login         string   `db:"login"`
	PasswordHash  string   `db:"password_hash"`
	IsBlocked     bool     `db:"is_blocked"`
	RoleID        uint64   `db:"role_id"`
	Role          Role     `db:"role"`
	PermissionIds []uint64 `db:"permissionIds"`
	ColumnIds     []uint64 `db:"columnIds"`
	PostIds       []uint64 `db:"postIds"`
}
