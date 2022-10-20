package feature

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

func (svc Service) saveFeature(ctx context.Context, f feature) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	f.ID = id
	return svc.store.saveFeature(ctx, f)
}
