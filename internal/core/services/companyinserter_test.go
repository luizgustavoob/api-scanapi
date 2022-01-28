package services_test

import (
	"errors"
	"testing"

	"github.com/api-scanapi/internal/core/entities"
	"github.com/api-scanapi/internal/core/ports/mocks"
	"github.com/api-scanapi/internal/core/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompanyInserter_AddCompany_WhenCompanyIsNil_ShouldReturnError(t *testing.T) {
	writer := new(mocks.CompanyWriter)

	service := services.NewCompanyInserter(writer)
	company, err := service.AddCompany(nil)

	assert.Nil(t, company)
	assert.NotNil(t, err)
}

func TestCompanyInserter_AddCompany_WhenInnerReturnError_ShouldReturnError(t *testing.T) {
	writer := new(mocks.CompanyWriter)
	writer.On("Insert", mock.Anything).Return(nil, errors.New("error"))

	service := services.NewCompanyInserter(writer)
	company, err := service.AddCompany(&entities.Company{})

	assert.Nil(t, company)
	assert.NotNil(t, err)
}

func TestCompanyInserter_AddCompany_WhenInnerReturnSuccess_ShouldReturnSuccess(t *testing.T) {
	expectedCompany := &entities.Company{
		ID:          "1",
		RazaoSocial: "Nome",
		CNPJ:        "cnpj",
		Cidade:      "cidade",
		UF:          "uf",
	}

	writer := new(mocks.CompanyWriter)
	writer.On("Insert", mock.Anything).Return(expectedCompany, nil)

	service := services.NewCompanyInserter(writer)
	company, err := service.AddCompany(&entities.Company{})

	assert.Equal(t, expectedCompany, company)
	assert.Nil(t, err)
}
