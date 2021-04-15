package storage

import "gitlab.neoway.com.br/companies/domain"

type CompanyStorageMock struct {
	InsertInvokedCount int

	InsertFn func(c *domain.Company) (*domain.Company, error)
}

func (self *CompanyStorageMock) Insert(c *domain.Company) (*domain.Company, error) {
	self.InsertInvokedCount++
	return self.InsertFn(c)
}
