package app

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	appService service.AppService
}

func NewController(appService service.AppService) *Controller {
	return &Controller{
		appService: appService,
	}
}

func (c *Controller) GetIsOnline(ctx echo.Context) error {
	isOnline := c.appService.IsOnline()
	err := ctx.JSON(http.StatusOK, isOnline)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
