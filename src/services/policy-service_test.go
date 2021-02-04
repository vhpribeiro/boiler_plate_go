package services

import (
	"errors"
	"testing"

	"controle_acesso_core.com/src/services/dtos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCasbinAdapter struct {
	mock.Mock
}

func (mock *MockCasbinAdapter) AddPolicy(role string, domain string, resource string, action string) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}

func TestShouldAddPolicyWithSuccess(t *testing.T) {
	//Arrange
	mockCasbinAdapter := new(MockCasbinAdapter)
	mockCasbinAdapter.On("AddPolicy").Return(true, nil)
	testService, _ := NewPolicyService(mockCasbinAdapter)
	role := "cadastro_lider"
	domain := "administrativo"
	resource := "cliente"
	action := "consultar"
	policyDto := dtos.PolicyDto{Role: role, Domain: domain, Resource: resource, Action: action}

	//Action
	result, _ := testService.AddPolicy(policyDto)

	//Assert
	mockCasbinAdapter.AssertExpectations(t)
	assert.Equal(t, true, result)
}

func TestShouldNotAddPolicy(t *testing.T) {
	//Arrange
	mockCasbinAdapter := new(MockCasbinAdapter)
	mockCasbinAdapter.On("AddPolicy").Return(false, nil)
	testService, _ := NewPolicyService(mockCasbinAdapter)
	role := "cadastro_lider"
	domain := "administrativo"
	resource := "cliente"
	action := "consultar"
	policyDto := dtos.PolicyDto{Role: role, Domain: domain, Resource: resource, Action: action}

	//Action
	result, _ := testService.AddPolicy(policyDto)

	//Assert
	mockCasbinAdapter.AssertExpectations(t)
	assert.Equal(t, false, result)
}

func TestShouldGetAErrorWhenAddPolicy(t *testing.T) {
	//Arrange
	mockCasbinAdapter := new(MockCasbinAdapter)
	mockCasbinAdapter.On("AddPolicy").Return(false, errors.New("Couldn't add policy"))
	testService, _ := NewPolicyService(mockCasbinAdapter)
	role := "cadastro_lider"
	domain := "administrativo"
	resource := "cliente"
	action := "consultar"
	policyDto := dtos.PolicyDto{Role: role, Domain: domain, Resource: resource, Action: action}

	//Action
	_, err := testService.AddPolicy(policyDto)

	//Assert
	mockCasbinAdapter.AssertExpectations(t)
	assert.NotNil(t, err)
}
