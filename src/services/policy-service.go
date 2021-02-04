package services

import (
	"controle_acesso_core.com/src/repositorys"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

type IPolicyService interface {
	AddPolicy(role string, domain string, resource string, action string) (bool, error)
}

type policyService struct {
	casbinRepo           repositorys.ICasbinRepository
	casbinMongoDbAdapter persist.BatchAdapter
	enforce              *casbin.Enforcer
}

func (p *policyService) AddPolicy(role string, domain string, resource string, action string) (bool, error) {
	return p.enforce.AddPolicy(role, domain, resource, action)
}

func NewPolicyService(casbinRepository repositorys.ICasbinRepository) (IPolicyService, error) {
	casbinContext, err := casbinRepository.GetTheAdapter()
	if err != nil {
		return nil, err
	}

	enforceConcrete, err := casbin.NewEnforcer("./configuration/casbin/casbin_rbac_with_domains_model.conf", casbinContext)
	if err != nil {
		return nil, err
	}

	return &policyService{
		casbinRepo:           casbinRepository,
		casbinMongoDbAdapter: casbinContext,
		enforce:              enforceConcrete,
	}, nil
}
