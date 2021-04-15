package client_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.neoway.com.br/companies/domain"
	"gitlab.neoway.com.br/companies/internal/infrastructure/client"
	httpclient "gitlab.neoway.com.br/companies/pkg/http"
)

const BRASIL_API_URL = "https://brasilapi.com.br/api"

func TestCompanyClient_IsValidCompany(t *testing.T) {

	t.Run("check company should return success code", func(t *testing.T) {
		httpClientMock := &httpclient.Mock{
			ResponseBody:   "",
			ResponseStatus: 200,
		}

		companyClient := client.NewCompanyClient(httpClientMock, BRASIL_API_URL)
		assert.NotNil(t, companyClient)

		var company = &domain.Company{
			RazaoSocial: "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA",
			CNPJ:        "05337875000105",
			Cidade:      "FLORIANOPOLIS",
			UF:          "SC",
		}

		isValid, err := companyClient.IsValidCompany(context.Background(), company)
		assert.NoError(t, err)
		assert.True(t, isValid)
	})

	t.Run("check company should return not_found_error code", func(t *testing.T) {
		httpClientMock := &httpclient.Mock{
			ResponseBody:   "",
			ResponseStatus: 404,
		}

		companyClient := client.NewCompanyClient(httpClientMock, BRASIL_API_URL)
		assert.NotNil(t, companyClient)

		var company = &domain.Company{
			RazaoSocial: "EMPRESA",
			CNPJ:        "00000000000000",
			Cidade:      "FLORIANOPOLIS",
			UF:          "SC",
		}

		isValid, err := companyClient.IsValidCompany(context.Background(), company)
		assert.NoError(t, err)
		assert.False(t, isValid)
	})

	t.Run("check company should return internal_server_error code", func(t *testing.T) {
		httpClientMock := &httpclient.Mock{
			ResponseBody:   "",
			ResponseStatus: 500,
		}

		companyClient := client.NewCompanyClient(httpClientMock, BRASIL_API_URL)
		assert.NotNil(t, companyClient)

		var company = &domain.Company{
			RazaoSocial: "EMPRESA",
			CNPJ:        "00000000000000",
			Cidade:      "FLORIANOPOLIS",
			UF:          "SC",
		}

		_, err := companyClient.IsValidCompany(context.Background(), company)
		assert.Error(t, err)
	})
}
