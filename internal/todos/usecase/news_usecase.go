package usecase

import (
	"context"

	"github.com/AliIsmoilov/golang_monolight/config"
	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/AliIsmoilov/golang_monolight/internal/todos"
	"github.com/AliIsmoilov/golang_monolight/pkg/logger"
	"github.com/AliIsmoilov/golang_monolight/pkg/utils"
	"github.com/google/uuid"
)

// News UseCase
type newsUC struct {
	cfg      *config.Config
	newsRepo todos.NewsRepository
	logger   logger.Logger
}

// News UseCase constructor
func NewNewsUseCase(cfg *config.Config, newsRepo todos.NewsRepository, logger logger.Logger) todos.NewsUseCase {
	return &newsUC{cfg: cfg, newsRepo: newsRepo, logger: logger}
}

// CreateNews
func (u *newsUC) Create(ctx context.Context, news *models.News) (*models.News, error) {
	return u.newsRepo.Create(ctx, news)
}

// Update news
func (u *newsUC) Update(ctx context.Context, news *models.News) (*models.News, error) {
	updatedNews, err := u.newsRepo.Update(ctx, news)
	if err != nil {
		return nil, err
	}

	return updatedNews, nil
}

// Delete news
func (u *newsUC) Delete(ctx context.Context, newsID uuid.UUID) error {

	if err := u.newsRepo.Delete(ctx, newsID); err != nil {
		return err
	}

	return nil
}

// Delete news
func (u *newsUC) SoftDelete(ctx context.Context, newsID uuid.UUID) error {

	if err := u.newsRepo.SoftDelete(ctx, newsID); err != nil {
		return err
	}

	return nil
}

// GetByID news
func (u *newsUC) GetByID(ctx context.Context, newID uuid.UUID) (*models.News, error) {

	return u.newsRepo.GetByID(ctx, newID)
}

// GetAll news
func (u *newsUC) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error) {
	return u.newsRepo.GetAll(ctx, title, query)
}
