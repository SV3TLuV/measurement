package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/column"
	"measurements-api/internal/server/http/middleware"
	"measurements-api/internal/service"
)

func InitColumnRoutes(group *echo.Group,
	columnService service.ColumnService,
	authMiddleware echo.MiddlewareFunc) {
	controller := column.NewController(columnService)
	baseGroup := group.Group("/columns")

	baseGroup.GET("", controller.Get, authMiddleware, middleware.AccessAdminMiddleware)
}
