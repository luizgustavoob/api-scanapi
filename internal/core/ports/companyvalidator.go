package ports

import (
	"context"

	"github.com/api-scanapi/internal/core/entities"
)

type CompanyValidator interface {
	CheckCompany(ctx context.Context, company *entities.Company) (bool, error)
}
