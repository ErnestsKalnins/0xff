package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func NewSaveEnvironmentHandler(strore Store) SaveEnvironmentHandler {
	return SaveEnvironmentHandler{store: strore}
}

type SaveEnvironmentHandler struct {
	store Store
}

func (h SaveEnvironmentHandler) Handle(ctx context.Context, e feature.Environment) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}
	e.ID = id

	now := time.Now().Unix()
	e.CreatedAt, e.UpdatedAt = now, now
	return h.store.SaveEnvironment(ctx, e)
}
