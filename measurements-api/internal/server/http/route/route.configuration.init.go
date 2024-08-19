package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/configuration"
	"measurements-api/internal/server/http/middleware"
	"measurements-api/internal/service"
)

func InitConfigurationRoutes(group *echo.Group,
	configurationService service.ConfigurationService,
	authMiddleware echo.MiddlewareFunc) {
	controller := configuration.NewController(configurationService)
	baseGroup := group.Group("/configuration")
	baseGroup.Use(authMiddleware, middleware.AccessAdminMiddleware)

	baseGroup.GET("", controller.Get)
	baseGroup.PUT("", controller.Put)
}
