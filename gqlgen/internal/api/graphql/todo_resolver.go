package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/exepirit/go-graphql/gqlgen/internal/api/graphql/dto"
	"github.com/exepirit/go-graphql/gqlgen/internal/api/graphql/gen"
	"github.com/exepirit/go-graphql/gqlgen/internal/models"
	"github.com/google/uuid"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input dto.NewTodo) (*dto.Todo, error) {
	todo := models.Todo{
		ID:   uuid.New(),
		Text: input.Text,
		Done: false,
	}

	err := r.todosRepository.Put(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &dto.Todo{
		ID:     todo.ID.String(),
		Text:   todo.Text,
		Done:   todo.Done,
		UserID: "",
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*dto.Todo, error) {
	todos, err := r.todosRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.Todo, len(todos))
	for i, todo := range todos {
		dtos[i] = &dto.Todo{
			ID:     todo.ID.String(),
			Text:   todo.Text,
			Done:   todo.Done,
			UserID: "",
		}
	}
	return dtos, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *dto.Todo) (*dto.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

// Todo returns gen.TodoResolver implementation.
func (r *Resolver) Todo() gen.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
