package internal

import (
	"github.com/api-scanapi/internal/adapters/clients"
	"github.com/api-scanapi/internal/adapters/handlers"
	"github.com/api-scanapi/internal/core"
	"github.com/api-scanapi/internal/repositories"
	"github.com/api-scanapi/internal/server"
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
