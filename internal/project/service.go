package project

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

func (svc Service) saveProject(ctx context.Context, p project) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	p.ID = id
	return svc.store.saveProject(ctx, p)
}
