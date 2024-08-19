package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/measurement"
	"measurements-api/internal/service"
)

func InitMeasurementRoutes(group *echo.Group,
	measurementService service.MeasurementService,
	authMiddleware echo.MiddlewareFunc) {
	controller := measurement.NewController(measurementService)
	baseGroup := group.Group("/measurements")

	private := baseGroup.Group("")
	private.Use(authMiddleware)

	private.GET("", controller.GetList)
	private.GET("/export", controller.Export)
}
