package ports

import "github.com/companies/internal/core/entities"

type CompanyWriter interface {
	Insert(company *entities.Company) (*entities.Company, error)
}
