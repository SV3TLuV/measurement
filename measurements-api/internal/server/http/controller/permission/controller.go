package permission

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/permission/converter"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	permissionService service.PermissionService
}

func NewController(permissionService service.PermissionService) *Controller {
	return &Controller{
		permissionService: permissionService,
	}
}

func (c *Controller) GetList(ctx echo.Context) error {
	context := ctx.Request().Context()
	permissions, err := c.permissionService.GetPermissions(context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter.ToPermissionViewsFromService(permissions)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
