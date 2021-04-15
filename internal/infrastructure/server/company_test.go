package http_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.neoway.com.br/companies/domain"
	"gitlab.neoway.com.br/companies/domain/company"
	httpInternal "gitlab.neoway.com.br/companies/internal/infrastructure/server"
)

func TestCompanyHandler_PostCompany(t *testing.T) {

	t.Run("should add company", func(t *testing.T) {
		serviceMock := &company.CompanyServiceMock{
			CheckCompanyFn: func(ctx context.Context, c *domain.Company) (bool, error) {
				return c.CNPJ != "" && c.UF != "", nil
			},
			AddCompanyFn: func(c *domain.Company) (*domain.Company, error) {
				c.ID = "abc123"
				return c, nil
			},
		}

		router := httpInternal.NewHandler(serviceMock)
		response := httptest.NewRecorder()
		body := []byte(`{"razaoSocial": "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA", "cnpj": "05337875000105", "cidade": "FLORIANOPOLIS", "uf": "SC"}`)

		req, _ := http.NewRequest("POST", "/v1/companies", bytes.NewReader(body))
		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusCreated, response.Code)

		body, err := ioutil.ReadAll(response.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		expectedBody := []byte(`{"id": "abc123", "razaoSocial": "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA", "cnpj": "05337875000105", "cidade": "FLORIANOPOLIS", "uf": "SC"}`)
		assert.JSONEq(t, string(expectedBody), string(body))
		assert.Equal(t, 1, serviceMock.CheckCompanyInvokedCount)
		assert.Equal(t, 1, serviceMock.AddCompanyInvokedCount)
	})
}
