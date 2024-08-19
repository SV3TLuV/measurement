package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/user"
	"measurements-api/internal/server/http/middleware"
	"measurements-api/internal/service"
)

func InitUserRoutes(group *echo.Group,
	userService service.UserService,
	authMiddleware echo.MiddlewareFunc) {
	controller := user.NewController(userService)
	baseGroup := group.Group("/users")
	baseGroup.Use(authMiddleware)

	baseGroup.GET("", controller.GetList)
	baseGroup.POST("", controller.Post,
		middleware.AccessAdminMiddleware)
	baseGroup.PUT("", controller.Put,
		middleware.AccessAdminMiddleware)

	baseGroup.GET("/me", controller.GetMe)

	baseGroup.GET("/:id/objects", controller.GetUserObjects,
		middleware.AccessAdminOrUserIdMiddleware)

	baseGroup.GET("/:id/columns", controller.GetUserColumns,
		middleware.AccessAdminOrUserIdMiddleware)

	baseGroup.GET("/:id/permissions", controller.GetUserPermissions,
		middleware.AccessAdminOrUserIdMiddleware)

	baseGroup.PUT("/change-password", controller.ChangePassword,
		middleware.AccessAdminMiddleware)
	baseGroup.PUT("/:id/ban", controller.Ban,
		middleware.AccessAdminMiddleware)
	baseGroup.PUT("/:id/unban", controller.Unban,
		middleware.AccessAdminMiddleware)
	baseGroup.DELETE("/:id", controller.Delete,
		middleware.AccessAdminMiddleware)
}
