package column

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/column/converter"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	columnService service.ColumnService
}

func NewController(columnService service.ColumnService) *Controller {
	return &Controller{
		columnService: columnService,
	}
}

func (c *Controller) Get(ctx echo.Context) error {
	context := ctx.Request().Context()
	columns, err := c.columnService.Get(context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter.ToColumnViewsFromService(columns)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
