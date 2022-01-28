package services

import (
	"errors"

	"github.com/api-scanapi/internal/core/entities"
	"github.com/api-scanapi/internal/core/ports"
)

type (
	insertService struct {
		companyWriter ports.CompanyWriter
	}
)

func NewCompanyInserter(companyWriter ports.CompanyWriter) ports.CompanyInserter {
	return &insertService{
		companyWriter: companyWriter,
	}
}

func (s *insertService) AddCompany(company *entities.Company) (*entities.Company, error) {
	if company == nil {
		return nil, errors.New("company is mandatory")
	}

	return s.companyWriter.Insert(company)
}
