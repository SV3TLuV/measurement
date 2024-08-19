package model

import "time"

type Measurement struct {
	ID                   uint64
	ObjectID             uint64
	ObjName              *string
	ObjNum               *string
	Created              time.Time
	Changed              *time.Time
	RealDateTime         *time.Time
	DateTime             *string
	Temp                 *float32
	Pressure             *float32
	WindDir              *int
	WindDirStr           *string
	WindSpeed            *float32
	Humid                *float32
	WaterVaporElasticity *float32
	AtmPhenom            *float32
	HumidInt             *float32
	TempInt              *float32

	V202917 *float32
	M202917 *float32
	Q202917 *string

	V202918 *float32
	M202918 *float32
	Q202918 *string

	V202919 *float32
	M202919 *float32
	Q202919 *string

	V202920 *float32
	M202920 *float32
	Q202920 *string

	V202921 *float32
	M202921 *float32
	Q202921 *string

	V202932 *float32
	M202932 *float32
	Q202932 *string

	V202935 *float32
	M202935 *float32
	Q202935 *string

	V202924 *float32
	M202924 *float32
	Q202924 *string

	V202925 *float32
	M202925 *float32
	Q202925 *string

	V203565 *float32
	M203565 *float32
	Q203565 *string

	V209190 *float32
	M209190 *float32
	Q209190 *string

	V203570 *float32
	M203570 *float32
	Q203570 *string

	V203551 *float32
	M203551 *float32
	Q203551 *string

	V202936 *float32
	M202936 *float32
	Q202936 *string

	V203569 *float32
	M203569 *float32
	Q203569 *string

	V203557 *float32
	M203557 *float32
	Q203557 *string

	V203568 *float32
	M203568 *float32
	Q203568 *string

	V203559 *float32
	M203559 *float32
	Q203559 *string

	V203577 *float32
	M203577 *float32
	Q203577 *string

	V211082 *float32
	M211082 *float32
	Q211082 *string

	V202931 *float32
	M202931 *float32
	Q202931 *string
}

func (m *Measurement) IsOld() bool {
	if m.Changed == nil {
		return false
	}

	threeWeeksAgoDate := time.Now().AddDate(0, 0, -21).UTC()
	return threeWeeksAgoDate.After(m.Changed.UTC())
}
