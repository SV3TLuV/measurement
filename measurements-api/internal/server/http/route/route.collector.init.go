package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/collector"
	"measurements-api/internal/server/http/middleware"
	"measurements-api/internal/service"
)

func InitCollectorRoutes(group *echo.Group,
	statisticService service.PollingStatisticService,
	collectorService service.CollectorService,
	authMiddleware echo.MiddlewareFunc) {
	controller := collector.NewController(statisticService, collectorService)
	baseGroup := group.Group("/collector")
	baseGroup.Use(authMiddleware, middleware.AccessAdminMiddleware)

	baseGroup.GET("/statistics", controller.GetStatistics)
	baseGroup.GET("/information", controller.GetInfo)
	baseGroup.GET("/state", controller.GetState)
}
