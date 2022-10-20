package feature

import (
	"github.com/google/uuid"
	"time"
)

type feature struct {
	ID            uuid.UUID `json:"id"`
	TechnicalName string    `json:"technicalName"`
	DisplayName   *string   `json:"displayName,omitempty"`
	Description   *string   `json:"description,omitempty"`
	Enabled       bool      `json:"enabled"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
