package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/exepirit/go-graphql/gqlgen/internal/api/graphql/gen"
	"github.com/exepirit/go-graphql/gqlgen/internal/models"
	"github.com/exepirit/go-graphql/gqlgen/internal/repository"
	"github.com/gin-gonic/gin"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todosRepository repository.EntityRepository[models.Todo]
}

func NewResolver(todos repository.EntityRepository[models.Todo]) *Resolver {
	return &Resolver{
		todosRepository: todos,
	}
}

func NewGraphqlEndpoints(resolvers *Resolver) *GraphqlEndpoints {
	gqlServer := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{
		Resolvers: resolvers,
	}))
	return &GraphqlEndpoints{
		Server: gqlServer,
	}
}

type GraphqlEndpoints struct {
	Server *handler.Server
}

func (handler GraphqlEndpoints) Bind(r gin.IRouter) {
	r.Any("", func(ctx *gin.Context) {
		handler.Server.ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func NewPlaygroundEndpoints() *PlaygroundEndpoints {
	return &PlaygroundEndpoints{}
}

type PlaygroundEndpoints struct{}

func (handler PlaygroundEndpoints) Bind(r gin.IRouter) {
	h := playground.Handler("ToDo Board", "/graphql")

	r.Any("", func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})
}
