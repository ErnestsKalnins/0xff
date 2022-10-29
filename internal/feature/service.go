package feature

import (
	"context"
	"fmt"
	"time"

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

	now := time.Now().Unix()
	f.CreatedAt, f.UpdatedAt = now, now
	return svc.store.saveFeature(ctx, f)
}

func (svc Service) setFeatureState(ctx context.Context, environmentID, featureID uuid.UUID, state featureState) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate id: %w", err)
	}

	return svc.store.saveFeatureState(ctx, writeState{
		ID:            id,
		FeatureID:     featureID,
		EnvironmentID: environmentID,
		State:         state,
		UpdatedAt:     time.Now().Unix(),
	})
}
