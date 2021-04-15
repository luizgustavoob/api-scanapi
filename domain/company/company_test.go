package company_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.neoway.com.br/companies/domain"
	"gitlab.neoway.com.br/companies/domain/company"
	"gitlab.neoway.com.br/companies/internal/infrastructure/client"
	"gitlab.neoway.com.br/companies/internal/infrastructure/storage"
)

var cpny = &domain.Company{
	RazaoSocial: "NEOWAY TECNOLOGIA INTEGRADA ASSESSORIA E NEGOCIOS SA",
	CNPJ:        "05337875000105",
	Cidade:      "FLORIANOPOLIS",
	UF:          "SC",
}

func TestCompanyService_AddCompany(t *testing.T) {

	myCompany := *cpny

	t.Run("should insert a company", func(t *testing.T) {
		storageMock := &storage.CompanyStorageMock{
			InsertFn: func(c *domain.Company) (*domain.Company, error) {
				c.ID = "abc123"
				return c, nil
			},
		}

		service := company.NewCompanyService(storageMock, nil)
		newCompany, err := service.AddCompany(&myCompany)

		assert.Equal(t, 1, storageMock.InsertInvokedCount)
		assert.NoError(t, err)
		assert.NotNil(t, newCompany)
	})
}

func TestCompanyService_CheckCompany(t *testing.T) {

	myCompany := *cpny

	t.Run("should check a valid company", func(t *testing.T) {
		clientMock := &client.CompanyClientMock{
			IsValidCompanyFn: func(ctx context.Context, c *domain.Company) (isValid bool, err error) {
				return c.CNPJ != "" && c.UF != "", nil
			},
		}

		service := company.NewCompanyService(nil, clientMock)
		isValid, err := service.CheckCompany(context.Background(), &myCompany)

		assert.Equal(t, 1, clientMock.IsValidCompanyInvokedCount)
		assert.NoError(t, err)
		assert.True(t, isValid)
	})
}
