package user

import (
	"github.com/labstack/echo/v4"
	converter3 "measurements-api/internal/server/http/controller/column/converter"
	converter4 "measurements-api/internal/server/http/controller/object/converter"
	converter2 "measurements-api/internal/server/http/controller/permission/converter"
	"measurements-api/internal/server/http/controller/user/converter"
	"measurements-api/internal/server/http/controller/user/model"
	"measurements-api/internal/server/http/utils"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	userService  service.UserService
	tokenService service.TokenService
}

func NewController(userService service.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (c *Controller) GetList(ctx echo.Context) error {
	var request model.FetchUserListRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	params := converter.ToGetUsersParamsFromRequest(&request)
	userList, err := c.userService.GetUsers(context, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	view := converter.ToUserPagedListViewFromService(userList)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) GetMe(ctx echo.Context) error {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	context := ctx.Request().Context()
	user, err := c.userService.GetUser(context, claims.ClientID)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	view := converter.ToUserViewFromService(user)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) GetUserObjects(ctx echo.Context) error {
	var request model.UserRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	objects, err := c.userService.GetUserObjects(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter4.ToObjectViewsFromService(objects)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) GetUserColumns(ctx echo.Context) error {
	var request model.UserRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	columns, err := c.userService.GetUserColumns(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter3.ToColumnViewsFromService(columns)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) GetUserPermissions(ctx echo.Context) error {
	var request model.UserRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	permissions, err := c.userService.GetUserPermissions(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter2.ToPermissionViewsFromService(permissions)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Post(ctx echo.Context) error {
	var request model.CreateUserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	user := converter.ToUserFromCreateRequest(&request)
	err := c.userService.Create(context, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Put(ctx echo.Context) error {
	var request model.UpdateUserRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	user := converter.ToUserFromUpdateRequest(&request)
	if err := c.userService.Update(context, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) ChangePassword(ctx echo.Context) error {
	var request model.ChangePasswordRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	err := c.userService.ChangePassword(context, request.UserID, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Ban(ctx echo.Context) error {
	var request model.UserRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	err := c.userService.Ban(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Unban(ctx echo.Context) error {
	var request model.UserRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	err := c.userService.Unban(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Delete(ctx echo.Context) error {
	var request model.UserRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	err := c.userService.Delete(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
