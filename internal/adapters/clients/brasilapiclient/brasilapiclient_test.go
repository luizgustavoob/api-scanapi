package brasilapiclient_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/companies/internal/adapters/clients/brasilapiclient"
	"github.com/companies/internal/adapters/clients/brasilapiclient/mocks"
	"github.com/companies/internal/core/entities"
	"github.com/stretchr/testify/assert"
)

func TestClient_IsValidCompany_WhenIsValid_ShouldReturnTrue(t *testing.T) {
	httpClient := new(mocks.HTTPClient)
	httpClient.On("Get", "https://brasilapi.com.br/api/cnpj/v1/20121850000155").Return(&http.Response{StatusCode: http.StatusOK}, nil)
	httpClient.On("Get", "https://brasilapi.com.br/api/ibge/uf/v1/SP").Return(&http.Response{StatusCode: http.StatusOK}, nil)

	company := &entities.Company{
		RazaoSocial: "MERCADO ENVIOS SERVICOS DE LOGISTICA LTDA",
		CNPJ:        "20121850000155",
		Cidade:      "OSASCO",
		UF:          "SP",
	}

	client := brasilapiclient.New(httpClient)
	isValid, err := client.IsValidCompany(context.TODO(), company)

	assert.Nil(t, err)
	assert.True(t, isValid)
}

func TestClient_IsValidCompany_WhenIsInvalid_ShouldReturnFalse(t *testing.T) {
	httpClient := new(mocks.HTTPClient)
	httpClient.On("Get", "https://brasilapi.com.br/api/cnpj/v1/20121850000155").Return(&http.Response{StatusCode: http.StatusOK}, nil)
	httpClient.On("Get", "https://brasilapi.com.br/api/ibge/uf/v1/AB").Return(&http.Response{StatusCode: http.StatusNotFound}, nil)

	company := &entities.Company{
		RazaoSocial: "MERCADO ENVIOS SERVICOS DE LOGISTICA LTDA",
		CNPJ:        "20121850000155",
		Cidade:      "OSASCO",
		UF:          "AB",
	}

	client := brasilapiclient.New(httpClient)
	isValid, err := client.IsValidCompany(context.TODO(), company)

	assert.Nil(t, err)
	assert.False(t, isValid)
}

func TestClient_IsValidCompany_WhenHttpClientReturnError_ShouldReturnError(t *testing.T) {
	httpClient := new(mocks.HTTPClient)
	httpClient.On("Get", "https://brasilapi.com.br/api/cnpj/v1/20121850000155").Return(&http.Response{StatusCode: http.StatusOK}, nil)
	httpClient.On("Get", "https://brasilapi.com.br/api/ibge/uf/v1/SP").Return(nil, errors.New("error"))

	company := &entities.Company{
		RazaoSocial: "MERCADO ENVIOS SERVICOS DE LOGISTICA LTDA",
		CNPJ:        "20121850000155",
		Cidade:      "OSASCO",
		UF:          "SP",
	}

	client := brasilapiclient.New(httpClient)
	isValid, err := client.IsValidCompany(context.TODO(), company)

	assert.NotNil(t, err)
	assert.False(t, isValid)
}
