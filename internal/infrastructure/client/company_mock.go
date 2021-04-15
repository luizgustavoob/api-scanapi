package client

import (
	"context"

	"gitlab.neoway.com.br/companies/domain"
)

type CompanyClientMock struct {
	IsValidCompanyInvokedCount int
	IsValidCompanyFn           func(ctx context.Context, c *domain.Company) (isValid bool, err error)
}

func (self *CompanyClientMock) IsValidCompany(ctx context.Context, c *domain.Company) (isValid bool, err error) {
	self.IsValidCompanyInvokedCount++
	return self.IsValidCompanyFn(ctx, c)
}
