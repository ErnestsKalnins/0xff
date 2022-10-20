package project

import "github.com/google/uuid"

type project struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
