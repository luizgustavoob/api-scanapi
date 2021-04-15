package company

import (
	"context"

	"gitlab.neoway.com.br/companies/domain"
)

type companyService struct {
	companyStorage domain.CompanyStorage
	companyClient  domain.CompanyClient
}

func NewCompanyService(storage domain.CompanyStorage, client domain.CompanyClient) *companyService {
	return &companyService{
		companyStorage: storage,
		companyClient:  client,
	}
}

func (self *companyService) CheckCompany(ctx context.Context, c *domain.Company) (bool, error) {
	return self.companyClient.IsValidCompany(ctx, c)
}

func (self *companyService) AddCompany(c *domain.Company) (*domain.Company, error) {
	return self.companyStorage.Insert(c)
}
