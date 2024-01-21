package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/AliIsmoilov/golang_monolight/internal/todos"
	"github.com/AliIsmoilov/golang_monolight/pkg/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// News Repository
type newsRepo struct {
	db *sqlx.DB
}

// ToDos Repository constructor
func NewNewsRepository(db *sqlx.DB) todos.NewsRepository {
	return &newsRepo{db: db}
}

// Create News
func (r *newsRepo) Create(ctx context.Context, new *models.News) (*models.News, error) {
	newUUID := uuid.New()
	c := &models.News{}
	createNews := `
		INSERT INTO news 
			(id, title, description, photo, published_by) 
		VALUES 
			($1, $2, $3, $4, $5) 
		RETURNING 
			id, title, description, photo, published_by, created_at`
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

// Update news
func (r *newsRepo) Update(ctx context.Context, new *models.News) (*models.News, error) {
	updateNews := `
		UPDATE news 
		SET 
			title = $1,
			description = $2,
			photo = $3,
			published_by = $4
		WHERE id = $5 
		RETURNING *`
	res := &models.News{}
	if err := r.db.
		QueryRowxContext(ctx, updateNews, new.Title, new.Description, new.Photo, new.PublishedBy, new.ID).
		StructScan(res); err != nil {
		return nil, errors.Wrap(err, "newsRepo.Update.QueryRowxContext")
	}

	return res, nil
}

// Delete news
func (r *newsRepo) Delete(ctx context.Context, newsID uuid.UUID) error {
	deleteNews := `DELETE FROM news WHERE id = $1`

	result, err := r.db.ExecContext(ctx, deleteNews, newsID)
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "newsRepo.Delete.rowsAffected")
	}

	return nil
}

// GetByID news
func (r *newsRepo) GetByID(ctx context.Context, newsId uuid.UUID) (*models.News, error) {
	getNewsByID := `
		SELECT id, title, description, photo, published_by, created_at
		FROM news
		WHERE id = $1`
	new := &models.News{}
	if err := r.db.GetContext(ctx, new, getNewsByID, newsId); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetByID.GetContext")
	}
	return new, nil
}

// GetAll news
func (r *newsRepo) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error) {
	var (
		totalCount    int
		getTotalCount = `SELECT COUNT(id) FROM news WHERE 1=1 AND deleted_at IS NULL`
		getAllNews    = `SELECT id, title, description, photo, published_by, created_at
							FROM news where 1=1 AND deleted_at IS NULL`
	)
	if title != "" {
		getTotalCount = fmt.Sprintf("%s%s", getTotalCount, " and title LIKE '%"+title+"%';")
		getAllNews = fmt.Sprintf("%s%s", getAllNews, " and title LIKE '%"+title+"%' ")
	}
	getAllNews += " ORDER BY created_at OFFSET $1 LIMIT $2;"
	if err := r.db.QueryRowContext(ctx, getTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.QueryRowContext")
	}

	if totalCount == 0 {
		return &models.NewsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			News:       make([]*models.News, 0),
		}, nil
	}

	rows, err := r.db.QueryxContext(ctx, getAllNews, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	newsList := make([]*models.News, 0, query.GetSize())
	for rows.Next() {
		new := &models.News{}
		if err = rows.StructScan(new); err != nil {
			return nil, errors.Wrap(err, "newsRepo.GetAll.StructScan")
		}
		newsList = append(newsList, new)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.rows.Err")
	}

	return &models.NewsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		News:       newsList,
	}, nil
}

// Update news
func (r *newsRepo) SoftDelete(ctx context.Context, newsID uuid.UUID) error {
	softDeleteNews := `UPDATE news SET deleted_at = now() WHERE id = $1`

	result, err := r.db.ExecContext(ctx, softDeleteNews, newsID)
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "newsRepo.Delete.rowsAffected")
	}

	return nil
}
