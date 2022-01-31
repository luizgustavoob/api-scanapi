package ports

import "github.com/companies/internal/core/entities"

type CompanyInserter interface {
	AddCompany(company *entities.Company) (*entities.Company, error)
}
