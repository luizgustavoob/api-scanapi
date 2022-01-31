package internal

import (
	"github.com/companies/internal/adapters/clients"
	"github.com/companies/internal/adapters/handlers"
	"github.com/companies/internal/core"
	"github.com/companies/internal/repositories"
	"github.com/companies/internal/server"
	"go.uber.org/fx"
)

func Serve() {
	app := fx.New(
		fx.Options(
			server.Module,
			clients.Module,
			handlers.Module,
			core.Module,
			repositories.Module,
		),
	)
	app.Run()
}
