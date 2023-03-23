package models

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID
	Name string
}

func (user User) GetID() uuid.UUID {
	return user.ID
}
