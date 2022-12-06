package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func NewSaveFeatureHandler(store Store) SaveFeatureHandler {
	return SaveFeatureHandler{store}
}

type SaveFeatureHandler struct {
	store Store
}

func (h SaveFeatureHandler) Handle(ctx context.Context, f feature.Feature) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}
	f.ID = id

	now := time.Now().Unix()
	f.CreatedAt, f.UpdatedAt = now, now
	return h.store.SaveFeature(ctx, f)
}
