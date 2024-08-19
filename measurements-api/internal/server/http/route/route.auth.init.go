package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/auth"
	"measurements-api/internal/service"
)

func InitAuthRoutes(group *echo.Group,
	authService service.AuthService,
	authMiddleware echo.MiddlewareFunc) {
	controller := auth.NewController(authService)
	baseGroup := group.Group("/auth")

	public := baseGroup.Group("")
	public.POST("/login", controller.Login)
	public.PUT("/refresh", controller.Refresh)

	private := baseGroup.Group("")
	private.Use(authMiddleware)

	private.POST("/logout", controller.Logout)
}
