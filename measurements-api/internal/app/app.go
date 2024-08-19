package app

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"measurements-api/internal/config"
	"measurements-api/internal/model"
	"measurements-api/internal/server/http/route"
	"measurements-api/internal/server/http/validator"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *echo.Echo
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	err := a.runScheduler()
	if err != nil {
		return err
	}
	return a.runHttpServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHttpServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	return config.Load()
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHttpServer(_ context.Context) error {
	a.httpServer = echo.New()

	a.httpServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: a.serviceProvider.Config().Jwt.Key,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.Claims)
		},
	})

	a.httpServer.Validator = validator.NewRequestValidator()

	group := a.httpServer.Group("/api")

	route.InitAppRoutes(group, a.serviceProvider.AppService())
	route.InitAuthRoutes(group, a.serviceProvider.AuthService(), jwtMiddleware)
	route.InitCollectorRoutes(group,
		a.serviceProvider.PollingStatisticService(),
		a.serviceProvider.CollectorService(), jwtMiddleware)
	route.InitColumnRoutes(group, a.serviceProvider.ColumnService(), jwtMiddleware)
	route.InitConfigurationRoutes(group, a.serviceProvider.ConfigurationService(), jwtMiddleware)
	route.InitMeasurementRoutes(group, a.serviceProvider.MeasurementService(), jwtMiddleware)
	route.InitObjectRoutes(group, a.serviceProvider.ObjectService(), jwtMiddleware)
	route.InitPermissionRoutes(group, a.serviceProvider.PermissionService(), jwtMiddleware)
	route.InitQualityRoutes(group, a.serviceProvider.QualityService(), jwtMiddleware)
	route.InitRouteRoles(group, a.serviceProvider.RoleService(), jwtMiddleware)
	route.InitUserRoutes(group, a.serviceProvider.UserService(), jwtMiddleware)

	return nil
}

func (a *App) runScheduler() error {
	schedulerService := a.serviceProvider.SchedulerService()
	err := schedulerService.Start()
	if err != nil {
		return errors.Wrap(err, "failed to start scheduler")
	}
	return nil
}

func (a *App) runHttpServer() error {
	err := a.httpServer.Start(a.serviceProvider.Config().Http.Addr())
	if err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}
