package ports

import (
	"context"

	"github.com/api-scanapi/internal/core/entities"
)

type CompanyClient interface {
	IsValidCompany(ctx context.Context, company *entities.Company) (bool, error)
}
