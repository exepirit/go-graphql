package graphql

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewResolver),
	fx.Provide(NewGraphqlEndpoints),
	fx.Provide(NewPlaygroundEndpoints),
)
