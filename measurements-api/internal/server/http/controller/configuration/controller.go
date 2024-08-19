package configuration

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/configuration/converter"
	"measurements-api/internal/server/http/controller/configuration/model"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	configurationService service.ConfigurationService
}

func NewController(configurationService service.ConfigurationService) *Controller {
	return &Controller{
		configurationService: configurationService,
	}
}

func (c *Controller) Get(ctx echo.Context) error {
	context := ctx.Request().Context()
	configuration, err := c.configurationService.Get(context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	view := converter.ToConfigurationViewFromService(configuration)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Put(ctx echo.Context) error {
	var request model.UpdateConfigurationRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	configuration := converter.ToConfigurationFromRequest(&request)
	err := c.configurationService.Update(context, configuration)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
