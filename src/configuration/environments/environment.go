package environments

import (
	"sync"

	"github.com/gosidekick/goconfig"
)

var (
	doOnceEnvironment sync.Once
	env               Environment
)

type Environment struct {
	Port             string `cfg:"PORT" cfgDefault:"8000"`
	ConnectionString string `cfg:"CONNECTION_STRING" cfgDefault:"postgresql://postgres:admin@localhost:5432/database?sslmode=disable"`
	ServiceName      string `cfg:"SERVICE_NAME" cfgDefault:"controle_acesso_core"`
}

func GetEnvironment() Environment {
	doOnceEnvironment.Do(func() {
		env = Environment{}
		err := goconfig.Parse(&env)
		if err != nil {
			println(err)
			return
		}
	})
	return env
}

func ReloadEnvironment() Environment {
	doOnceEnvironment = *new(sync.Once)
	return GetEnvironment()
}
