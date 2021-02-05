package main

import (
	"controle_acesso_core.com/src/controllers"
	http_adapter "controle_acesso_core.com/src/http"
	"controle_acesso_core.com/src/repositorys"
	"controle_acesso_core.com/src/services"
)

func main() {

	//Instanciar os repositorys
	casbinRepository := repositorys.NewCasbinPostgressRepository()

	//Instanciar os serviços
	casbinAdapterService, err := services.NewCasbinAdapterService(casbinRepository)
	if err != nil {
		panic(err)
	}
	policyService, err := services.NewPolicyService(casbinAdapterService)
	if err != nil {
		panic(err)
	}

	policyController := controllers.NewPolicyController(policyService)
	healthCheckController := controllers.NewHealtCheckController()

	handler := http_adapter.NewHandler(healthCheckController, policyController)

	err = handler.Start()
	if err != nil {
		panic(err)
	}
}
