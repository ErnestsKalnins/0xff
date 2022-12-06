package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func NewSaveProjectHandler(store Store) SaveProjectHandler {
	return SaveProjectHandler{store: store}
}

type SaveProjectHandler struct {
	store Store
}

func (h SaveProjectHandler) Handle(ctx context.Context, p feature.Project) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}
	p.ID = id

	now := time.Now().Unix()
	p.CreatedAt, p.UpdatedAt = now, now
	return h.store.SaveProject(ctx, p)
}
