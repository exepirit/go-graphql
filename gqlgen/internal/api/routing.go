package api

import (
	"github.com/exepirit/go-graphql/gqlgen/internal/api/graphql"
	"github.com/exepirit/go-graphql/gqlgen/pkg/routing"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewAPI),
)

type API routing.Bindable

func NewAPI(
	graphql *graphql.GraphqlEndpoints,
	graphqlPlayground *graphql.PlaygroundEndpoints,
) API {
	return routing.Union(
		routing.Route("/graphql", graphql),
		routing.Route("/playground", graphqlPlayground),
	)
}
