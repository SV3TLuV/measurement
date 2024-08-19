package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/role"
	"measurements-api/internal/server/http/middleware"
	"measurements-api/internal/service"
)

func InitRouteRoles(group *echo.Group,
	roleService service.RoleService,
	authMiddleware echo.MiddlewareFunc) {
	controller := role.NewController(roleService)
	baseGroup := group.Group("/roles")

	baseGroup.GET("", controller.GetList, authMiddleware,
		middleware.AccessAdminMiddleware)
}
