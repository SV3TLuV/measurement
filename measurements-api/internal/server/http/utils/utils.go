package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"measurements-api/internal/model"
	"net/http"
)

func GetClaims(ctx echo.Context) (*model.Claims, error) {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized)
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized)
	}

	return claims, nil
}

func GetUserId(ctx echo.Context) uint64 {
	claims, _ := GetClaims(ctx)
	return claims.ClientID
}
