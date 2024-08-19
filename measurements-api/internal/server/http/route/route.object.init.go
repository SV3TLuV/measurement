package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/object"
	"measurements-api/internal/server/http/middleware"
	"measurements-api/internal/service"
)

func InitObjectRoutes(group *echo.Group,
	objectService service.ObjectService,
	authMiddleware echo.MiddlewareFunc) {
	controller := object.NewController(objectService)
	baseGroup := group.Group("/objects")
	baseGroup.Use(authMiddleware)

	baseGroup.GET("", controller.GetList)
	baseGroup.GET("/posts/:id", controller.GetPost)
	baseGroup.GET("/search-new", controller.SearchNew,
		middleware.AccessAdminMiddleware)
	baseGroup.PUT("/:id/enable", controller.Enable,
		middleware.AccessAdminMiddleware)
	baseGroup.PUT("/:id/disable", controller.Disable,
		middleware.AccessAdminMiddleware)
}
