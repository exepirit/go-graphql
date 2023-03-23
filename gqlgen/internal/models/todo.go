package models

import "github.com/google/uuid"

type Todo struct {
	ID uuid.UUID
	Text string
	Done bool
}

func (todo Todo) GetID() uuid.UUID {
	return todo.ID
}
