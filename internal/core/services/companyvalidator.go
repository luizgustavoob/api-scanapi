package services

import (
	"context"
	"errors"

	"github.com/companies/internal/core/entities"
	"github.com/companies/internal/core/ports"
)

type (
	validatorService struct {
		companyClient ports.CompanyClient
	}
)

func NewCompanyValidator(companyClient ports.CompanyClient) ports.CompanyValidator {
	return &validatorService{
		companyClient: companyClient,
	}
}

func (s *validatorService) CheckCompany(ctx context.Context, company *entities.Company) (bool, error) {
	if company == nil {
		return false, errors.New("company is mandatory")
	}

	return s.companyClient.IsValidCompany(ctx, company)
}
