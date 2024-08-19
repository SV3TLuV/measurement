package entities

type Quality struct {
	ID       uint64 `db:"quality_id"`
	Priority int    `db:"priority"`
	Title    string `db:"title"`
	Color    string `db:"color"`
	Caption  string `db:"caption"`
}
