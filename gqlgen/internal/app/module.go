package app

import (
	"github.com/exepirit/go-graphql/gqlgen/gql"
	"github.com/exepirit/go-graphql/gqlgen/internal/api"
	"github.com/exepirit/go-graphql/gqlgen/internal/config"
	"github.com/exepirit/go-graphql/gqlgen/internal/infrastructure"
	"github.com/exepirit/go-graphql/gqlgen/internal/repository/memory"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	infrastructure.Module,
	memory.Module,
	gql.Module,
	api.Module,
	fx.Invoke(bootstrap),
)
