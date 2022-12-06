package feature

import (
	"fmt"

	"github.com/google/uuid"
)

type Environment struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"projectId"`
	Name      string    `json:"name"`
	CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
}

type ErrEnvironmentNotFound struct {
	ID uuid.UUID
}

func (e ErrEnvironmentNotFound) Error() string {
	return fmt.Sprintf("could not find environment by id %s", e.ID)
}
