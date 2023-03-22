package repository

import (
	"context"

	"github.com/google/uuid"
)

type EntityRepository[T Entity] interface {
	GetAll(ctx context.Context) ([]T, error)
	Get(ctx context.Context, id uuid.UUID) (T, error)
	Put(ctx context.Context, entity T) error
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Entity interface {
	ID() uuid.UUID
}
