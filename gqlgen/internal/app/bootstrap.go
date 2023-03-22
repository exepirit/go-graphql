package app

import (
	"context"

	"github.com/exepirit/go-graphql/gqlgen/internal/api"
	"github.com/exepirit/go-graphql/gqlgen/internal/config"
	"github.com/exepirit/go-graphql/gqlgen/internal/infrastructure"
	"go.uber.org/fx"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	cfg config.Config,
	server infrastructure.Server,
	api api.API,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Bind(api)

			go func(server infrastructure.Server) {
				if err := server.ListenAndServe(); err != nil {
					panic(err)
				}
			}(server)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
