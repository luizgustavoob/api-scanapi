package brasilapiclient_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/api-scanapi/internal/adapters/clients/brasilapiclient"
	"github.com/api-scanapi/internal/adapters/clients/brasilapiclient/mocks"
	"github.com/api-scanapi/internal/core/entities"
	"github.com/stretchr/testify/assert"
)

func TestClient_IsValidCompany_WhenIsValid_ShouldReturnTrue(t *testing.T) {
	httpClient := new(mocks.HTTPClient)
	httpClient.On("Get", "https://brasilapi.com.br/api/cnpj/v1/05337875000105").Return(&http.Response{StatusCode: http.StatusOK}, nil)
	httpClient.On("Get", "https://brasilapi.com.br/api/ibge/uf/v1/SC").Return(&http.Response{StatusCode: http.StatusOK}, nil)

	company := &entities.Company{
		RazaoSocial: "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA",
		CNPJ:        "05337875000105",
		Cidade:      "FLORIANOPOLIS",
		UF:          "SC",
	}

	client := brasilapiclient.New(httpClient)
	isValid, err := client.IsValidCompany(context.TODO(), company)

	assert.Nil(t, err)
	assert.True(t, isValid)
}

func TestClient_IsValidCompany_WhenIsInvalid_ShouldReturnFalse(t *testing.T) {
	httpClient := new(mocks.HTTPClient)
	httpClient.On("Get", "https://brasilapi.com.br/api/cnpj/v1/05337875000105").Return(&http.Response{StatusCode: http.StatusOK}, nil)
	httpClient.On("Get", "https://brasilapi.com.br/api/ibge/uf/v1/AB").Return(&http.Response{StatusCode: http.StatusNotFound}, nil)

	company := &entities.Company{
		RazaoSocial: "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA",
		CNPJ:        "05337875000105",
		Cidade:      "FLORIANOPOLIS",
		UF:          "AB",
	}

	client := brasilapiclient.New(httpClient)
	isValid, err := client.IsValidCompany(context.TODO(), company)

	assert.Nil(t, err)
	assert.False(t, isValid)
}

func TestClient_IsValidCompany_WhenHttpClientReturnError_ShouldReturnError(t *testing.T) {
	httpClient := new(mocks.HTTPClient)
	httpClient.On("Get", "https://brasilapi.com.br/api/cnpj/v1/05337875000105").Return(&http.Response{StatusCode: http.StatusOK}, nil)
	httpClient.On("Get", "https://brasilapi.com.br/api/ibge/uf/v1/SC").Return(nil, errors.New("error"))

	company := &entities.Company{
		RazaoSocial: "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA",
		CNPJ:        "05337875000105",
		Cidade:      "FLORIANOPOLIS",
		UF:          "SC",
	}

	client := brasilapiclient.New(httpClient)
	isValid, err := client.IsValidCompany(context.TODO(), company)

	assert.NotNil(t, err)
	assert.False(t, isValid)
}
