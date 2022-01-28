package ports

import "github.com/api-scanapi/internal/core/entities"

type CompanyWriter interface {
	Insert(company *entities.Company) (*entities.Company, error)
}
