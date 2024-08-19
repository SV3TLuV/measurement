package model

type ObjectTypeKey uint64

const (
	LaboratoryKey ObjectTypeKey = iota + 1
	CityKey
	PostKey
)

type ObjectType struct {
	ID    uint64
	Title string
}
