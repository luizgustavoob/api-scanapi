package ports

import (
	"context"

	"github.com/companies/internal/core/entities"
)

type CompanyClient interface {
	IsValidCompany(ctx context.Context, company *entities.Company) (bool, error)
}
