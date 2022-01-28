package server

import (
	"github.com/api-scanapi/internal/server/httpstarter"
	"github.com/api-scanapi/internal/server/postgresql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	httpstarter.Module,
	postgresql.Module,
)
