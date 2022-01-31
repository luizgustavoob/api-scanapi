package repositories

import (
	"database/sql"

	"github.com/companies/internal/core/ports"
	"github.com/companies/internal/repositories/companiespostgresql"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	newCompanyWriter,
)

func newCompanyWriter(db *sql.DB) ports.CompanyWriter {
	return companiespostgresql.NewWriter(db)
}
