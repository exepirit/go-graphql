package logic

import (
	"github.com/exepirit/go-graphql/gqlgen/internal/models"
	"github.com/exepirit/go-graphql/gqlgen/pkg/namegen"
	"github.com/google/uuid"
)

func CreateUser() models.User {
	return models.User{
		ID:   uuid.New(),
		Name: namegen.NewName(),
	}
}
