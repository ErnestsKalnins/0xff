package feature

import (
	"fmt"

	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
}

type ErrProjectNotFound struct {
	ID uuid.UUID
}

func (e ErrProjectNotFound) Error() string {
	return fmt.Sprintf("could not find project by id %s", e.ID)
}
