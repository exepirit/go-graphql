package memory

import (
	"context"
	"errors"

	"github.com/exepirit/go-graphql/gqlgen/internal/repository"
	"github.com/google/uuid"
)

func NewEntityRepository[T repository.Entity]() *EntityRepository[T] {
	return &EntityRepository[T]{
		entities: make(map[uuid.UUID]T),
	}
}

type EntityRepository[T repository.Entity] struct {
	entities map[uuid.UUID]T
}

func (repo EntityRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	entities := make([]T, 0, len(repo.entities))
	for _, entity := range repo.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (repo EntityRepository[T]) Get(ctx context.Context, id uuid.UUID) (T, error) {
	entity, ok := repo.entities[id]
	if !ok {
		return *(new(T)), errors.New("not found")
	}
	return entity, nil
}

func (repo *EntityRepository[T]) Put(ctx context.Context, entity T) error {
	_, ok := repo.entities[entity.GetID()]
	if ok {
		return errors.New("already exists")
	}
	repo.entities[entity.GetID()] = entity
	return nil
}

func (repo *EntityRepository[T]) Update(ctx context.Context, entity T) error {
	_, ok := repo.entities[entity.GetID()]
	if !ok {
		return errors.New("not found")
	}
	repo.entities[entity.GetID()] = entity
	return nil
}

func (repo *EntityRepository[T]) Delete(ctx context.Context, id uuid.UUID) error {
	_, ok := repo.entities[id]
	if !ok {
		return errors.New("not found")
	}
	delete(repo.entities, id)
	return nil
}
