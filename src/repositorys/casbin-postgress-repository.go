package repositorys

import (
	"controle_acesso_core.com/src/utils"
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2/persist"
)

type casbinPostgressRepository struct {
	envLoader utils.IEnvLoader
}

func (repo *casbinPostgressRepository) GetTheAdapter() (persist.BatchAdapter, error) {
	connectionString, err := repo.envLoader.GetTheEnvVariable("CONNECTION_STRING")
	if err != nil {
		return nil, err
	}

	adapter, err := pgadapter.NewAdapter(connectionString)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}

func NewCasbinPostgressRepository(envLoader utils.IEnvLoader) ICasbinRepository {
	return &casbinPostgressRepository{
		envLoader: envLoader,
	}
}
