//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package todos

import (
	"context"

	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/AliIsmoilov/golang_monolight/pkg/utils"
	"github.com/google/uuid"
)

// ToDos repository interface
type BlogRepository interface {
	Create(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Update(ctx context.Context, todo *models.Blog) (*models.Blog, error)
	Delete(ctx context.Context, todoID uuid.UUID) error
	GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error)
	GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogsList, error)

	CreateNews(ctx context.Context, new *models.News) (*models.News, error)
}
