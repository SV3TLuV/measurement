package model

type ColumnView struct {
	ID         uint64  `json:"id"`
	Title      string  `json:"title"`
	ShortTitle string  `json:"shortTitle"`
	Formula    *string `json:"formula"`
	ObjField   string  `json:"objField"`
	Code       *string `json:"code"`
}
