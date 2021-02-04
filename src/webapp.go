package main

import (
	"controle_acesso_core.com/src/controllers"
	http_adapter "controle_acesso_core.com/src/http"
	"controle_acesso_core.com/src/repositorys"
	"controle_acesso_core.com/src/services"
	"controle_acesso_core.com/src/utils"
)

func main() {
	//Instanciar utils
	envLoader := utils.NewEnvLoader()

	//Instanciar os repositorys
	casbinRepository := repositorys.NewCasbinPostgressRepository(envLoader)

	//Instanciar os serviços
	policyService, err := services.NewPolicyService(casbinRepository)
	if err != nil {
		panic(err)
	}

	policyController := controllers.NewPolicyController(policyService)
	healthCheckController := controllers.NewHealtCheckController()

	handler := http_adapter.NewHandler(healthCheckController, policyController, envLoader)

	err = handler.Start()
	if err != nil {
		panic(err)
	}
}
