package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/permission"
	"measurements-api/internal/service"
)

func InitPermissionRoutes(group *echo.Group,
	permissionService service.PermissionService,
	authMiddleware echo.MiddlewareFunc) {
	controller := permission.NewController(permissionService)
	baseGroup := group.Group("/permissions")

	baseGroup.GET("", controller.GetList, authMiddleware)
}
