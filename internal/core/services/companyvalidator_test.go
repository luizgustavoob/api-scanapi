package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/companies/internal/core/entities"
	"github.com/companies/internal/core/ports/mocks"
	"github.com/companies/internal/core/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCompanyValidator_CheckCompany_WhenCompanyIsNil_ShouldReturnError(t *testing.T) {
	client := new(mocks.CompanyClient)

	service := services.NewCompanyValidator(client)

	isValid, err := service.CheckCompany(context.TODO(), nil)

	assert.False(t, isValid)
	assert.NotNil(t, err)
}

func TestCompanyValidator_CheckCompany_WhenInnerReturnError_ShouldReturnError(t *testing.T) {
	client := new(mocks.CompanyClient)
	client.On("IsValidCompany", mock.Anything, mock.Anything).Return(false, errors.New("error"))

	service := services.NewCompanyValidator(client)

	isValid, err := service.CheckCompany(context.TODO(), &entities.Company{})

	assert.False(t, isValid)
	assert.NotNil(t, err)
}

func TestCompanyValidator_CheckCompany_WhenInnerReturnSuccess_ShouldReturnSuccess(t *testing.T) {
	client := new(mocks.CompanyClient)
	client.On("IsValidCompany", mock.Anything, mock.Anything).Return(true, nil)

	service := services.NewCompanyValidator(client)

	isValid, err := service.CheckCompany(context.TODO(), &entities.Company{})

	assert.True(t, isValid)
	assert.Nil(t, err)
}
