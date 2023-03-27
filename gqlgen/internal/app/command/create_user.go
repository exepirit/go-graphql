package command

import (
	"context"
	"fmt"

	"github.com/exepirit/go-graphql/gqlgen/internal/models"
	"github.com/exepirit/go-graphql/gqlgen/internal/repository"
	"github.com/exepirit/go-graphql/gqlgen/pkg/namegen"
	"github.com/google/uuid"
)

func NewCreateUserCommand(usersRepo repository.EntityRepository[models.User]) *CreateUserCommand {
	return &CreateUserCommand{usersRepo: usersRepo}
}

type CreateUserArgs struct{}

type CreateUserCommand struct{
	usersRepo repository.EntityRepository[models.User]
}

func (cmd CreateUserCommand) Execute(ctx context.Context, args CreateUserArgs) (uuid.UUID, error) {
	user := models.User{
		ID:   uuid.New(),
		Name: namegen.NewName(),
	}

	if err := cmd.usersRepo.Put(ctx, user); err != nil {
		return uuid.Nil, fmt.Errorf("failed to store new user: %w", err)
	}

	return user.ID, nil
}
