package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/auth/converter"
	"measurements-api/internal/server/http/controller/auth/model"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	authService service.AuthService
}

func NewController(authService service.AuthService) *Controller {
	return &Controller{
		authService: authService,
	}
}

func (c *Controller) Login(ctx echo.Context) error {
	var request model.LoginRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	command := converter.ToLoginDataFromRequest(&request)
	result, err := c.authService.Login(context, command)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := converter.ToAuthResultViewFromService(result)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return nil
}

func (c *Controller) Logout(ctx echo.Context) error {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	context := ctx.Request().Context()
	if err := c.authService.Logout(context, token); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *Controller) Refresh(ctx echo.Context) error {
	var request model.RefreshRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	result, err := c.authService.Refresh(context, request.RefreshToken)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := converter.ToAuthResultViewFromService(result)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return nil
}
