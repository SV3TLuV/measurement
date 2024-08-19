package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/quality"
	"measurements-api/internal/service"
)

func InitQualityRoutes(group *echo.Group,
	qualityService service.QualityService,
	authMiddleware echo.MiddlewareFunc) {
	controller := quality.NewController(qualityService)
	baseGroup := group.Group("/qualities")

	baseGroup.GET("", controller.GetList, authMiddleware)
}
