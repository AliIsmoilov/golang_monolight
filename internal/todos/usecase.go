//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package todos

import (
	"context"

	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/AliIsmoilov/golang_monolight/pkg/utils"
	"github.com/google/uuid"
)

// blogs use case
type UseCase interface {
	Create(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Update(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Delete(ctx context.Context, blogID uuid.UUID) error
	GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error)
	GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogsList, error)
}

// News use case
type NewsUseCase interface {
	Create(ctx context.Context, News *models.News) (*models.News, error)
	Update(ctx context.Context, News *models.News) (*models.News, error)
	Delete(ctx context.Context, NewsID uuid.UUID) error
	SoftDelete(ctx context.Context, NewsID uuid.UUID) error
	GetByID(ctx context.Context, NewsID uuid.UUID) (*models.News, error)
	GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error)
}
