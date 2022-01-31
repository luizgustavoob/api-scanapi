package server

import (
	"github.com/companies/internal/server/httpstarter"
	"github.com/companies/internal/server/postgresql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	httpstarter.Module,
	postgresql.Module,
)
