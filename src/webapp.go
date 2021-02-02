package main

import (
	"controle_acesso_core.com/src/controllers"
	http_adapter "controle_acesso_core.com/src/http"
)

func main() {
	healthCheckController := controllers.NewHealtCheckController()

	handler := http_adapter.NewHandler(healthCheckController)

	err := handler.Start()
	if err != nil {
		panic(err)
	}
}
