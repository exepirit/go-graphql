package command

import (
	"context"

	"github.com/google/uuid"
)

type Command[TA any] interface {
	Execute(ctx context.Context, args TA) (uuid.UUID, error)
}
