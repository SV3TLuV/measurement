package object

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	model2 "measurements-api/internal/model"
	"measurements-api/internal/server/http/controller/object/converter"
	"measurements-api/internal/server/http/controller/object/model"
	converter2 "measurements-api/internal/server/http/controller/user/converter"
	"measurements-api/internal/server/http/utils"
	"measurements-api/internal/service"
	"net/http"
)

type Controller struct {
	objectService service.ObjectService
}

func NewController(objectService service.ObjectService) *Controller {
	return &Controller{
		objectService: objectService,
	}
}

func (c *Controller) GetList(ctx echo.Context) error {
	var request model.FetchObjectListQuery
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	params := converter2.ToGetObjectsParamsFromQuery(&request)
	objects, err := c.objectService.GetObjects(context, params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter.ToObjectViewsFromService(objects)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) GetPost(ctx echo.Context) error {
	var request model.ObjectRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	userId := utils.GetUserId(ctx)
	object, err := c.objectService.GetPost(context, userId, request.ID)
	if errors.Is(err, model2.NotFound) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	view := converter.ToObjectViewFromService(object)
	if err = ctx.JSON(http.StatusOK, view); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) SearchNew(ctx echo.Context) error {
	context := ctx.Request().Context()
	posts, err := c.objectService.SearchNew(context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	views := converter.ToObjectViewsFromService(posts)
	if err = ctx.JSON(http.StatusOK, views); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Enable(ctx echo.Context) error {
	var request model.ObjectRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	err := c.objectService.Enable(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}

func (c *Controller) Disable(ctx echo.Context) error {
	var request model.ObjectRequestWithID
	if err := ctx.Bind(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err := ctx.Validate(&request); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	context := ctx.Request().Context()
	err := c.objectService.Disable(context, request.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unknown error")
	}

	return nil
}
