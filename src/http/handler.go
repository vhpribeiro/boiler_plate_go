package http

import (
	"fmt"

	"boiler_plate.com/src/configuration/environments"
	"boiler_plate.com/src/controllers"
	"boiler_plate.com/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IHandler interface {
	Start() error
}

type apiHandler struct {
	echo        *echo.Echo
	healthCheck controllers.IHealthCheckController
	envLoader   utils.IEnvLoader
}

func (handler *apiHandler) Start() error {
	env := environments.GetEnvironment()

	handler.echo.Use(middleware.Logger())
	handler.echo.Use(middleware.Recover())
	api := handler.echo.Group("/api")

	healthCheck := api.Group("/health")
	healthCheck.GET("", handler.healthCheck.GetHealthCheck)

	return handler.echo.Start(fmt.Sprintf("%s:%s", "0.0.0.0", env.Port))
}

func NewHandler(
	healthCheckController controllers.IHealthCheckController) IHandler {
	e := echo.New()
	return &apiHandler{
		echo:        e,
		healthCheck: healthCheckController,
	}
}
