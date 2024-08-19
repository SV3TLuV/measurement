package quality

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/quality/converter"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	qualityService service.QualityService
}

func NewController(qualityService service.QualityService) *Controller {
	return &Controller{
		qualityService: qualityService,
	}
}

func (c *Controller) GetList(ctx echo.Context) error {
	context := ctx.Request().Context()
	qualities, err := c.qualityService.GetQualities(context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter.ToQualityViewsFromService(qualities)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
