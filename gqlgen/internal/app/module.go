package app

import (
	"github.com/exepirit/go-graphql/gqlgen/internal/api"
	"github.com/exepirit/go-graphql/gqlgen/internal/api/graphql"
	"github.com/exepirit/go-graphql/gqlgen/internal/app/command"
	"github.com/exepirit/go-graphql/gqlgen/internal/config"
	"github.com/exepirit/go-graphql/gqlgen/internal/infrastructure"
	"github.com/exepirit/go-graphql/gqlgen/internal/repository/memory"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	infrastructure.Module,
	memory.Module,
	command.Module,
	graphql.Module,
	api.Module,
	fx.Invoke(bootstrap),
)
