package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/exepirit/go-graphql/gqlgen/internal/api/graphql/gen"
	"github.com/exepirit/go-graphql/gqlgen/internal/models"
	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input gen.NewTodo) (*gen.Todo, error) {
	todo := models.Todo{
		ID:   uuid.New(),
		Text: input.Text,
		Done: false,
	}

	err := r.todosRepository.Put(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &gen.Todo{
		ID:   todo.ID.String(),
		Text: todo.Text,
		Done: todo.Done,
		User: nil,
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*gen.Todo, error) {
	todos, err := r.todosRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]*gen.Todo, len(todos))
	for i, todo := range todos {
		dtos[i] = &gen.Todo{
			ID:   todo.ID.String(),
			Text: todo.Text,
			Done: todo.Done,
			User: nil,
		}
	}
	return dtos, nil
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
