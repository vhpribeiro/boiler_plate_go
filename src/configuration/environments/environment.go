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
	Port                 string `cfg:"PORT" cfgDefault:"8000"`
	PostgresHost         string `cfg:"POSTGRES_HOST" cfgDefault:"localhost"`
	PostgresPort         string `cfg:"POSTGRES_PORT" cfgDefault:"5432"`
	PostgresUser         string `cfg:"POSTGRES_USER" cfgDefault:"postgres"`
	PostgresPassword     string `cfg:"POSTGRES_PASSWORD" cfgDefault:"admin"`
	PostgresDatabaseName string `cfg:"POSTGRES_DATABASE_NAME" cfgDefault:"postgres_users"`
	RedisHost            string `cfg:"REDIS_HOST" cfgDefault:"localhost"`
	RedisPort            string `cfg:"REDIS_Port" cfgDefault:"6379"`
	ServiceName          string `cfg:"SERVICE_NAME" cfgDefault:"boiler_plate"`
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
