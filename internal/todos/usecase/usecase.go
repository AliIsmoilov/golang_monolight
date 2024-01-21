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

// ToDos UseCase
type todosUC struct {
	cfg       *config.Config
	blogsRepo todos.BlogRepository
	logger    logger.Logger
}

// ToDos UseCase constructor
func NewToDosUseCase(cfg *config.Config, blogsRepo todos.BlogRepository, logger logger.Logger) todos.UseCase {
	return &todosUC{cfg: cfg, blogsRepo: blogsRepo, logger: logger}
}

// Create todo
func (u *todosUC) Create(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	return u.blogsRepo.Create(ctx, blog)
}

// Update todo
func (u *todosUC) Update(ctx context.Context, todo *models.Blog) (*models.Blog, error) {
	updatedToDo, err := u.blogsRepo.Update(ctx, todo)
	if err != nil {
		return nil, err
	}

	return updatedToDo, nil
}

// Delete todo
func (u *todosUC) Delete(ctx context.Context, todoID uuid.UUID) error {

	if err := u.blogsRepo.Delete(ctx, todoID); err != nil {
		return err
	}

	return nil
}

// GetByID todo
func (u *todosUC) GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error) {

	return u.blogsRepo.GetByID(ctx, blogID)
}

// GetAll todos
func (u *todosUC) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogsList, error) {
	return u.blogsRepo.GetAll(ctx, title, query)
}
