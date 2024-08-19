package model

type AuthResultView struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
