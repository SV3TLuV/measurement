package model

import (
	"measurements-api/internal/model"
	"time"
)

type GetMeasurementsParams struct {
	UserID   uint64
	ObjectID *uint64
	Period   *model.Period
	Start    *time.Time
	End      *time.Time
	Page     uint
	PageSize uint
}
