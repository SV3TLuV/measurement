package collector

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/collector/converter"
	"measurements-api/internal/server/http/controller/collector/model"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	statisticService service.PollingStatisticService
	collectorService service.CollectorService
}

func NewController(statisticService service.PollingStatisticService,
	collectorService service.CollectorService) *Controller {
	return &Controller{
		statisticService: statisticService,
		collectorService: collectorService,
	}
}

func (c *Controller) GetStatistics(ctx echo.Context) error {
	var request model.FetchPollingStatisticQuery
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	params := converter.ToGetPollingStatisticsParamsFromRequest(&request)
	statistics, err := c.statisticService.GetStatistics(context, params)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	views := converter.ToPollingStatisticViewsFromService(statistics)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return nil
}

func (c *Controller) GetInfo(ctx echo.Context) error {
	context := ctx.Request().Context()
	information, err := c.collectorService.GetInfo(context)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := converter.ToInformationViewFromService(information)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return nil
}

func (c *Controller) GetState(ctx echo.Context) error {
	context := ctx.Request().Context()
	state, err := c.collectorService.GetState(context)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	view := converter.ToControllerStateViewFromService(state)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return nil
}
