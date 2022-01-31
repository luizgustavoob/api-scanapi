package handlers

import (
	"github.com/companies/internal/adapters/handlers/companieshandler"
	"github.com/companies/internal/adapters/handlers/pinghandler"
	"github.com/companies/internal/core/ports"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	newPingHandler,
	newCompaniesHandler,
)

type (
	HandlerResult struct {
		fx.Out
		Handler ports.Handler `group:"handlers"`
	}
)

func newPingHandler() HandlerResult {
	return HandlerResult{
		Handler: pinghandler.New(),
	}
}

func newCompaniesHandler(
	validator ports.CompanyValidator,
	inserter ports.CompanyInserter,
) HandlerResult {
	return HandlerResult{
		Handler: companieshandler.New(validator, inserter),
	}
}
