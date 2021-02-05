package http

import (
	"fmt"

	"controle_acesso_core.com/src/configuration/environments"
	"controle_acesso_core.com/src/controllers"
	"controle_acesso_core.com/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qiangxue/go-env"
)

type IHandler interface {
	Start() error
}

type apiHandler struct {
	echo        *echo.Echo
	healthCheck controllers.IHealthCheckController
	policys     controllers.IPolicyController
	envLoader   utils.IEnvLoader
}

func (handler *apiHandler) Start() error {
	var cfg environments.Environment
	if err := env.Load(&cfg); err != nil {
		panic(err)
	}

	handler.echo.Use(middleware.Logger())
	handler.echo.Use(middleware.Recover())
	api := handler.echo.Group("/api")

	healthCheck := api.Group("/healthcheck")
	healthCheck.GET("", handler.healthCheck.GetHealthCheck)

	policys := api.Group("/policies")
	policys.POST("", handler.policys.AddPolicy)

	return handler.echo.Start(fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Port))
}

func NewHandler(
	healthCheckController controllers.IHealthCheckController,
	policyController controllers.IPolicyController) IHandler {
	e := echo.New()
	return &apiHandler{
		echo:        e,
		healthCheck: healthCheckController,
		policys:     policyController,
	}
}
