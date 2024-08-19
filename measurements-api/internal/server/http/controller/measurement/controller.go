package measurement

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/measurement/converter"
	"measurements-api/internal/server/http/controller/measurement/model"
	"measurements-api/internal/server/http/utils"
	"measurements-api/internal/service"
	"net/http"
	"os"
)

type Controller struct {
	measurementService service.MeasurementService
}

func NewController(measurementService service.MeasurementService) *Controller {
	return &Controller{
		measurementService: measurementService,
	}
}

func (c *Controller) GetList(ctx echo.Context) error {
	var request model.FetchMeasurementListQuery
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	params := converter.ToGetMeasurementsParamsFromQuery(&request)
	params.UserID = utils.GetUserId(ctx)
	list, err := c.measurementService.GetMeasurements(context, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter.ToMeasurementPagedListViewFromService(list)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Export(ctx echo.Context) error {
	const fileName = "measurement.csv"
	var request model.ExportMeasurementListQuery
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	params := converter.ToGetMeasurementParamsFromExportQuery(request)
	params.UserID = utils.GetUserId(ctx)
	context := ctx.Request().Context()
	bytes, err := c.measurementService.Export(context, *params, request.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	contentType := converter.ToContentTypeFromExportFormat(request.Format)
	if len(contentType) == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	ctx.Response().Header().Set(echo.HeaderContentType, contentType)

	tmpFile, err := os.CreateTemp("", fileName)
	if err != nil {
		return err
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(tmpFile.Name())

	if _, err := tmpFile.Write(bytes); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	if err := tmpFile.Close(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	if err = ctx.Attachment(tmpFile.Name(), fileName); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
