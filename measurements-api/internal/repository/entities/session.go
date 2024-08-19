package entities

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           uuid.UUID  `db:"session_id"`
	UserID       uint64     `db:"user_id"`
	Created      time.Time  `db:"created"`
	Updated      *time.Time `db:"updated"`
	RefreshToken string     `db:"refresh_token"`
}
