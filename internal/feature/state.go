package feature

import (
	"time"

	"github.com/google/uuid"
)

type state struct {
	ID            uuid.UUID `json:"id"`
	FeatureID     uuid.UUID `json:"featureId"`
	EnvironmentID uuid.UUID `json:"environmentId"`
	Enabled       bool      `json:"enabled"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
