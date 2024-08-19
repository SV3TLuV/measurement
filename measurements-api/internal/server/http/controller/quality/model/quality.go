package model

type QualityView struct {
	ID       uint64 `json:"id"`
	Priority int    `json:"priority"`
	Title    string `json:"title"`
	Color    string `json:"color"`
	Caption  string `json:"caption"`
}
