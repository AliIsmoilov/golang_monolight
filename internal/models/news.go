package models

import (
	"time"

	"github.com/google/uuid"
)

// Blog Swagger model
type NewsSwagger struct {
	Title       string    `json:"title" db:"title" validate:"required,gte=3"`
	Description string    `json:"description" db:"description"`
	Photo       uuid.UUID `json:"photo" db:"photo"`
	PublishedBy uuid.UUID `json:"published_by" db:"published_by"`
}

type News struct {
	ID          uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Title       string    `json:"title" db:"title" validate:"required,gte=3"`
	Description string    `json:"description" db:"description"`
	Photo       uuid.UUID `json:"photo" db:"photo"`
	PublishedBy uuid.UUID `json:"published_by" db:"published_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
