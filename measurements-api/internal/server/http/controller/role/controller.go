package role

import (
	"github.com/labstack/echo/v4"
	"measurements-api/internal/server/http/controller/role/converter"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	roleService service.RoleService
}

func NewController(roleService service.RoleService) *Controller {
	return &Controller{
		roleService: roleService,
	}
}

func (c *Controller) GetList(ctx echo.Context) error {
	context := ctx.Request().Context()
	roles, err := c.roleService.GetRoles(context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter.ToRoleViewsFromService(roles)

	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
