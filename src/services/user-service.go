package services

import (
	"fmt"

	"controle_acesso_core.com/src/repositorys"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

type IUserService interface {
	AddRoleForUserInDomain(user string, domain string, role string) (bool, error)
	CheckIfUserHasPermission(user string, domain string, resource string, action string) (bool, error)
}

type userService struct {
	casbinRepository       repositorys.ICasbinRepository
	casbinPostgreesAdapter persist.BatchAdapter
	enforce                *casbin.Enforcer
}

func (userService *userService) AddRoleForUserInDomain(user string, domain string, role string) (bool, error) {
	result, err := userService.enforce.AddRoleForUserInDomain(user, role, domain)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (userService *userService) CheckIfUserHasPermission(user string, domain string, resource string, action string) (bool, error) {
	fmt.Printf("\nuser: %v\ndomain: %v\nresource: %v\naction: %v\n", user, domain, resource, action)
	result, err := userService.enforce.Enforce(user, domain, resource, action)
	if err != nil {
		return false, err
	}

	return result, nil
}

func NewUserService(casbinRepository repositorys.ICasbinRepository) (IUserService, error) {
	casbinContext, err := casbinRepository.GetTheAdapter()
	if err != nil {
		return nil, err
	}

	enforceConcrete, err := casbin.NewEnforcer("./configuration/casbin_rbac_with_domains_model.conf", casbinContext)
	if err != nil {
		return nil, err
	}

	return &userService{
		casbinRepository:       casbinRepository,
		casbinPostgreesAdapter: casbinContext,
		enforce:                enforceConcrete,
	}, nil
}
