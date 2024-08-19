package route

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/app"
	"measurements-api/internal/service"
)

func InitAppRoutes(group *echo.Group, appService service.AppService) {
	controller := app.NewController(appService)

	group.GET("/online", controller.GetIsOnline)
}
