package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ErnestsKalnins/0xff/feature"
)

func NewSaveFeatureStateHandler(store Store) SaveFeatureStateHandler {
	return SaveFeatureStateHandler{store: store}
}

type SaveFeatureStateHandler struct {
	store Store
}

type SaveFeatureStateRequest struct {
	EnvironmentID uuid.UUID
	FeatureID     uuid.UUID
	State         feature.State
}

func (h SaveFeatureStateHandler) Handle(ctx context.Context, req SaveFeatureStateRequest) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	return h.store.SaveFeatureState(ctx, feature.EnvironmentFeature{
		ID:            id,
		FeatureID:     req.FeatureID,
		EnvironmentID: req.EnvironmentID,
		State:         req.State,
		UpdatedAt:     time.Now().Unix(),
	})
}
