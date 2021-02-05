package repositorys

import (
	"controle_acesso_core.com/src/configuration/environments"
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2/persist"
)

type casbinPostgressRepository struct{}

func (repo *casbinPostgressRepository) GetTheAdapter() (persist.BatchAdapter, error) {

	env := environments.GetEnvironment()

	adapter, err := pgadapter.NewAdapter(env.ConnectionString)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}

func NewCasbinPostgressRepository() ICasbinRepository {
	return &casbinPostgressRepository{}
}
