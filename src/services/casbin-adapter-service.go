package services

import (
	"controle_acesso_core.com/src/repositorys"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

type ICasbinAdapterService interface {
	AddPolicy(role string, domain string, resource string, action string) (bool, error)
}

type casbinAdapterService struct {
	casbinRepo           repositorys.ICasbinRepository
	casbinMongoDbAdapter persist.BatchAdapter
	enforce              *casbin.Enforcer
}

func (c *casbinAdapterService) AddPolicy(role string, domain string, resource string, action string) (bool, error) {
	return c.enforce.AddPolicy(role, domain, resource, action)
}

func NewCasbinAdapterService(casbinRepository repositorys.ICasbinRepository) (ICasbinAdapterService, error) {
	casbinContext, err := casbinRepository.GetTheAdapter()
	if err != nil {
		return nil, err
	}

	enforceConcrete, err := casbin.NewEnforcer("./configuration/casbin/casbin_rbac_with_domains_model.conf", casbinContext)
	if err != nil {
		return nil, err
	}

	return &casbinAdapterService{
		casbinRepo:           casbinRepository,
		casbinMongoDbAdapter: casbinContext,
		enforce:              enforceConcrete,
	}, nil
}
