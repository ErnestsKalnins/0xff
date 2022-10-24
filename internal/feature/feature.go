package feature

import (
	"github.com/google/uuid"
)

type feature struct {
	ID            uuid.UUID `json:"id"`
	ProjectID     uuid.UUID `json:"projectId"`
	TechnicalName string    `json:"technicalName"`
	DisplayName   *string   `json:"displayName,omitempty"`
	Description   *string   `json:"description,omitempty"`
	CreatedAt     int64     `json:"createdAt"`
	UpdatedAt     int64     `json:"updatedAt"`
}
