package repositorys

import (
	"controle_acesso_core.com/src/configuration/environments"
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2/persist"
	"github.com/qiangxue/go-env"
)

type casbinPostgressRepository struct{}

func (repo *casbinPostgressRepository) GetTheAdapter() (persist.BatchAdapter, error) {

	var cfg environments.Environment
	if err := env.Load(&cfg); err != nil {
		panic(err)
	}

	adapter, err := pgadapter.NewAdapter(cfg.ConnectionString)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}

func NewCasbinPostgressRepository() ICasbinRepository {
	return &casbinPostgressRepository{}
}
