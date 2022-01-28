package core

import (
	"github.com/api-scanapi/internal/core/ports"
	"github.com/api-scanapi/internal/core/services"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	newCompanyInserter,
	newCompanyValidator,
)

func newCompanyInserter(companyWriter ports.CompanyWriter) ports.CompanyInserter {
	return services.NewCompanyInserter(companyWriter)
}

func newCompanyValidator(companyClient ports.CompanyClient) ports.CompanyValidator {
	return services.NewCompanyValidator(companyClient)
}
