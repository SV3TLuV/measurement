package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type LoginData struct {
	Login    string
	Password string
}

type AuthResult struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	ClientID uint64    `json:"client_id"`
	SID      uuid.UUID `json:"sid"`
	Role     string    `json:"roles"`
	jwt.RegisteredClaims
}
