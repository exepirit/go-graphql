package memory

import (
	"github.com/exepirit/go-graphql/gqlgen/internal/models"
	"github.com/exepirit/go-graphql/gqlgen/internal/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(func() repository.EntityRepository[models.Todo] {
		return NewEntityRepository[models.Todo]()
	}),
)
