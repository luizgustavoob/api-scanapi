package company

import (
	"context"

	"gitlab.neoway.com.br/companies/domain"
)

type CompanyServiceMock struct {
	CheckCompanyInvokedCount int
	AddCompanyInvokedCount   int

	CheckCompanyFn func(ctx context.Context, c *domain.Company) (bool, error)
	AddCompanyFn   func(c *domain.Company) (*domain.Company, error)
}

func (self *CompanyServiceMock) CheckCompany(ctx context.Context, c *domain.Company) (bool, error) {
	self.CheckCompanyInvokedCount++
	return self.CheckCompanyFn(ctx, c)
}

func (self *CompanyServiceMock) AddCompany(c *domain.Company) (*domain.Company, error) {
	self.AddCompanyInvokedCount++
	return self.AddCompanyFn(c)
}
