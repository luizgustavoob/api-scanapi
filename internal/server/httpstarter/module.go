package httpstarter

import (
	"context"
	"net/http"
	"time"

	"github.com/companies/internal/core/ports"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		Register,
	),
	fx.Invoke(
		StartServer,
	),
)

type (
	HandlerParams struct {
		fx.In
		Handlers []ports.Handler `group:"handlers"`
	}
)

func Register(params HandlerParams) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.Default()

	for _, h := range params.Handlers {
		handler.Handle(h.GetHttpMethod(), h.GetRelativePath(), gin.WrapH(h))
	}

	return handler
}

func StartServer(lc fx.Lifecycle, router http.Handler) {
	srv := &http.Server{
		Addr:         ":9998",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
