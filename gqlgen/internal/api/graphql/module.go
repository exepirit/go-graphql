package graphql

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGraphqlEndpoints),
	fx.Provide(NewPlaygroundEndpoints),
)
