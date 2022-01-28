package ports

import "github.com/api-scanapi/internal/core/entities"

type CompanyInserter interface {
	AddCompany(company *entities.Company) (*entities.Company, error)
}
