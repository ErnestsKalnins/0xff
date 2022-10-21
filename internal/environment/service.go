package environment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func NewService(store Store) Service {
	return Service{store: store}
}

type Service struct {
	store Store
}

func (svc Service) saveEnvironment(ctx context.Context, e environment) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	e.ID = id
	return svc.store.saveEnvironment(ctx, e)
}
