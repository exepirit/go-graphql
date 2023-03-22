package api

import (
	"github.com/exepirit/go-graphql/gqlgen/gql"
	"github.com/exepirit/go-graphql/gqlgen/pkg/routing"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewAPI),
)

type API routing.Bindable

func NewAPI(
	graphql gql.GraphqlHandler,
) API {
	return routing.Union(
		routing.Route("/graphql", graphql),
	)
}
