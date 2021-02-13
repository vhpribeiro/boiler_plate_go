package main

import (
	"boiler_plate.com/src/controllers"
	http_adapter "boiler_plate.com/src/http"
)

func main() {

	//Instanciar os repositorys

	//Instanciar os serviços

	//Instanciar controllers
	healthCheckController := controllers.NewHealtCheckController()

	handler := http_adapter.NewHandler(healthCheckController)

	err := handler.Start()
	if err != nil {
		panic(err)
	}
}
