package repository

import (
	"context"

	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Create todo
func (r *blogsRepo) CreateNews(ctx context.Context, new *models.News) (*models.News, error) {
	newUUID := uuid.New()
	c := &models.News{}
	createNews := `
		INSERT INTO news 
			(id, title, description, photo, published_by) 
		VALUES 
			($1, $2, $3, $4, $5) 
		RETURNING *`
	if err := r.db.QueryRowxContext(
		ctx,
		createNews,
		newUUID,
		&new.Title,
		&new.Description,
		&new.Photo,
		&new.PublishedBy,
	).StructScan(c); err != nil {
		return nil, errors.Wrap(err, "newsRepo.Create.StructScan")
	}

	return c, nil
}
