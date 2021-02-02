package repository

import (
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2/persist"
)

type casbinPostgressRepository struct {
	connectionString string
}

func (repo *casbinPostgressRepository) GetTheAdapter() (persist.BatchAdapter, error) {
	adapter, err := pgadapter.NewAdapter(repo.connectionString)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}

func NewCasbinPostgressRepository() ICasbinRepository {
	return &casbinPostgressRepository{
		connectionString: "postgresql://username:password@postgres:5432/database?sslmode=disable",
	}
}
