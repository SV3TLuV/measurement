package model

import (
	"measurements-api/internal/model"
	"time"
)

type MeasurementView struct {
	ID                   uint64     `json:"id"`
	ObjectID             uint64     `json:"objectId"`
	ObjName              *string    `json:"objName"`
	ObjNum               *string    `json:"objNum"`
	Created              time.Time  `json:"created"`
	Changed              *time.Time `json:"changed"`
	RealDateTime         *time.Time `json:"realDateTime"`
	DateTime             *string    `json:"dateTime"`
	Temp                 *float32   `json:"temp"`
	Pressure             *float32   `json:"pressure"`
	WindDir              *int       `json:"windDir"`
	WindDirStr           *string    `json:"windDirStr"`
	WindSpeed            *float32   `json:"windSpeed"`
	Humid                *float32   `json:"humidity"`
	WaterVaporElasticity *float32   `json:"waterVaporElasticity"`
	AtmPhenom            *float32   `json:"atmPhenom"`
	HumidInt             *float32   `json:"humidInt"`
	TempInt              *float32   `json:"tempInt"`

	V202917 *float32 `json:"v202917"`
	M202917 *float32 `json:"m202917"`
	Q202917 *string  `json:"q202917"`

	V202918 *float32 `json:"v202918"`
	M202918 *float32 `json:"m202918"`
	Q202918 *string  `json:"q202918"`

	V202919 *float32 `json:"v202919"`
	M202919 *float32 `json:"m202919"`
	Q202919 *string  `json:"q202919"`

	V202920 *float32 `json:"v202920"`
	M202920 *float32 `json:"m202920"`
	Q202920 *string  `json:"q202920"`

	V202921 *float32 `json:"v202921"`
	M202921 *float32 `json:"m202921"`
	Q202921 *string  `json:"q202921"`

	V202932 *float32 `json:"v202932"`
	M202932 *float32 `json:"m202932"`
	Q202932 *string  `json:"q202932"`

	V202935 *float32 `json:"v202935"`
	M202935 *float32 `json:"m202935"`
	Q202935 *string  `json:"q202935"`

	V202924 *float32 `json:"v202924"`
	M202924 *float32 `json:"m202924"`
	Q202924 *string  `json:"q202924"`

	V202925 *float32 `json:"v202925"`
	M202925 *float32 `json:"m202925"`
	Q202925 *string  `json:"q202925"`

	V203565 *float32 `json:"v203565"`
	M203565 *float32 `json:"m203565"`
	Q203565 *string  `json:"q203565"`

	V209190 *float32 `json:"v209190"`
	M209190 *float32 `json:"m209190"`
	Q209190 *string  `json:"q209190"`

	V203570 *float32 `json:"v203570"`
	M203570 *float32 `json:"m203570"`
	Q203570 *string  `json:"q203570"`

	V203551 *float32 `json:"v203551"`
	M203551 *float32 `json:"m203551"`
	Q203551 *string  `json:"q203551"`

	V202936 *float32 `json:"v202936"`
	M202936 *float32 `json:"m202936"`
	Q202936 *string  `json:"q202936"`

	V203569 *float32 `json:"v203569"`
	M203569 *float32 `json:"m203569"`
	Q203569 *string  `json:"q203569"`

	V203557 *float32 `json:"v203557"`
	M203557 *float32 `json:"m203557"`
	Q203557 *string  `json:"q203557"`

	V203568 *float32 `json:"v203568"`
	M203568 *float32 `json:"m203568"`
	Q203568 *string  `json:"q203568"`

	V203559 *float32 `json:"v203559"`
	M203559 *float32 `json:"m203559"`
	Q203559 *string  `json:"q203559"`

	V203577 *float32 `json:"v203577"`
	M203577 *float32 `json:"m203577"`
	Q203577 *string  `json:"q203577"`

	V211082 *float32 `json:"v211082"`
	M211082 *float32 `json:"m211082"`
	Q211082 *string  `json:"q211082"`

	V202931 *float32 `json:"v202931"`
	M202931 *float32 `json:"m202931"`
	Q202931 *string  `json:"q202931"`
}

type FetchMeasurementListQuery struct {
	ObjectID *uint64       `query:"objectId"`
	Period   *model.Period `query:"period"`
	Start    *time.Time    `query:"start"`
	End      *time.Time    `query:"end"`
	Page     uint          `query:"page" validate:"gte=1"`
	PageSize uint          `query:"pageSize" validate:"gte=1"`
}

type ExportMeasurementListQuery struct {
	ObjectID *uint64            `query:"objectId"`
	Period   *model.Period      `query:"period"`
	Start    *time.Time         `query:"start"`
	End      *time.Time         `query:"end"`
	Page     uint               `query:"page" validate:"gte=0"`
	PageSize uint               `query:"pageSize" validate:"gte=0"`
	Format   model.ExportFormat `query:"format" validate:"required"`
}
