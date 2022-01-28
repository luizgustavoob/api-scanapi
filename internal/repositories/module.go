package repositories

import (
	"database/sql"

	"github.com/api-scanapi/internal/core/ports"
	"github.com/api-scanapi/internal/repositories/companiespostgresql"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	newCompanyWriter,
)

func newCompanyWriter(db *sql.DB) ports.CompanyWriter {
	return companiespostgresql.NewWriter(db)
}
