package services

import "controle_acesso_core.com/src/services/dtos"

type IPolicyService interface {
	AddPolicy(policyDto dtos.PolicyDto) (bool, error)
}

type policyService struct {
	casbinAdapter ICasbinAdapterService
}

func (p *policyService) AddPolicy(policyDto dtos.PolicyDto) (bool, error) {
	return p.casbinAdapter.AddPolicy(policyDto.Role, policyDto.Domain, policyDto.Resource, policyDto.Action)
}

func NewPolicyService(casbinAdapter ICasbinAdapterService) (IPolicyService, error) {
	return &policyService{
		casbinAdapter: casbinAdapter,
	}, nil
}
