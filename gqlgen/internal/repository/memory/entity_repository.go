package memory

import (
	"context"
	"errors"
	"reflect"

	"github.com/exepirit/go-graphql/gqlgen/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewEntityRepository[T repository.Entity](log *zap.Logger) *EntityRepository[T] {
	typeName := reflect.TypeOf(*(new(T))).Name()
	log = log.With(zap.String("module", "repository"), zap.String("entity", typeName))

	return &EntityRepository[T]{
		log: log,
		entities: make(map[uuid.UUID]T),
	}
}

type EntityRepository[T repository.Entity] struct {
	log *zap.Logger
	entities map[uuid.UUID]T
}

func (repo EntityRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	entities := make([]T, 0, len(repo.entities))
	for _, entity := range repo.entities {
		entities = append(entities, entity)
	}
	repo.log.Debug("entities scanned from hashmap")
	return entities, nil
}

func (repo EntityRepository[T]) Get(ctx context.Context, id uuid.UUID) (T, error) {
	entity, ok := repo.entities[id]
	if !ok {
		repo.log.Debug("entity not found", zap.String("op", "get"))
		return *(new(T)), repository.ErrNotFound
	}
	repo.log.Debug("entity found in hashmap")
	return entity, nil
}

func (repo *EntityRepository[T]) Put(ctx context.Context, entity T) error {
	_, ok := repo.entities[entity.GetID()]
	if ok {
		repo.log.Debug("entity already exists", zap.String("op", "put"))
		return errors.New("already exists")
	}
	repo.entities[entity.GetID()] = entity
	repo.log.Debug("new entity stored")
	return nil
}

func (repo *EntityRepository[T]) Update(ctx context.Context, entity T) error {
	_, ok := repo.entities[entity.GetID()]
	if !ok {
		repo.log.Debug("entity not found", zap.String("op", "update"))
		return repository.ErrNotFound
	}
	repo.entities[entity.GetID()] = entity
	repo.log.Debug("entity updated")
	return nil
}

func (repo *EntityRepository[T]) Delete(ctx context.Context, id uuid.UUID) error {
	_, ok := repo.entities[id]
	if !ok {
		repo.log.Debug("entity not found")
		return repository.ErrNotFound
	}
	delete(repo.entities, id)
	repo.log.Debug("entity deleted")
	return nil
}
