package model

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           uuid.UUID
	UserID       uint64
	Created      time.Time
	Updated      *time.Time
	RefreshToken string
}
