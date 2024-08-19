package entities

type Role struct {
	ID    uint64 `db:"role_id"`
	Title string `db:"title"`
	Name  string `db:"name"`
}
