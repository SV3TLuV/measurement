package middleware

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/utils"
	"net/http"
	"slices"
	"strconv"
)

var AccessAdminMiddleware = NewAccessByRoleMiddleware([]string{"Admin"})
var AccessAdminOrUserIdMiddleware = NewAccessByRoleOrUserIdMiddleware([]string{"Admin"})

func NewAccessByRoleMiddleware(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			claims, err := utils.GetClaims(ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			if !slices.Contains(roles, claims.Role) {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			return next(ctx)
		}
	}
}

func NewAccessByRoleOrUserIdMiddleware(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			claims, err := utils.GetClaims(ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			userId := strconv.FormatUint(claims.ClientID, 10)
			requestId := ctx.Param("id")

			if slices.Contains(roles, claims.Role) {
				return next(ctx)
			}

			if userId == requestId {
				return next(ctx)
			}

			return echo.NewHTTPError(http.StatusForbidden)
		}
	}
}
