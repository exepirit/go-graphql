package memory

import (
	"github.com/exepirit/go-graphql/gqlgen/internal/models"
	"github.com/exepirit/go-graphql/gqlgen/internal/repository"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(func(log *zap.Logger) repository.EntityRepository[models.Todo] {
		return NewEntityRepository[models.Todo](log)
	}),
	fx.Provide(func(log *zap.Logger) repository.EntityRepository[models.User] {
		return NewEntityRepository[models.User](log)
	}),
)
