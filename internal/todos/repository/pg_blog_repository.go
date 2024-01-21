package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/AliIsmoilov/golang_monolight/internal/todos"
	"github.com/AliIsmoilov/golang_monolight/pkg/utils"
)

// Blog Repository
type blogsRepo struct {
	db *sqlx.DB
}

// ToDos Repository constructor
func NewToDosRepository(db *sqlx.DB) todos.BlogRepository {
	return &blogsRepo{db: db}
}

// Create todo
func (r *blogsRepo) Create(ctx context.Context, todo *models.Blog) (*models.Blog, error) {
	newUUID := uuid.New()
	c := &models.Blog{}
	createBlog := `INSERT INTO blogs (id, title) VALUES ($1, $2) RETURNING *`
	if err := r.db.QueryRowxContext(
		ctx,
		createBlog,
		newUUID,
		&todo.Title,
	).StructScan(c); err != nil {
		return nil, errors.Wrap(err, "blogRepo.Create.StructScan")
	}

	return c, nil
}

// Update blog
func (r *blogsRepo) Update(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	updateBlog := `UPDATE blogs SET title = $1 WHERE id = $2 RETURNING *`
	res := &models.Blog{}
	if err := r.db.QueryRowxContext(ctx, updateBlog, blog.Title, blog.ID).StructScan(res); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.Update.QueryRowxContext")
	}

	return res, nil
}

// Delete blog
func (r *blogsRepo) Delete(ctx context.Context, blogID uuid.UUID) error {
	deleteBlog := `DELETE FROM blogs WHERE id = $1`

	result, err := r.db.ExecContext(ctx, deleteBlog, blogID)
	if err != nil {
		return errors.Wrap(err, "blogsRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "blogsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "blogsRepo.Delete.rowsAffected")
	}

	return nil
}

// GetByID blog
func (r *blogsRepo) GetByID(ctx context.Context, blogId uuid.UUID) (*models.Blog, error) {
	getBlogByID := `SELECT id, title, created_at
	FROM blogs
	WHERE id = $1`
	blog := &models.Blog{}
	if err := r.db.GetContext(ctx, blog, getBlogByID, blogId); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetByID.GetContext")
	}
	return blog, nil
}

// GetAll ToDos
func (r *blogsRepo) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogsList, error) {
	var (
		totalCount    int
		getTotalCount = `SELECT COUNT(id) FROM blogs WHERE 1=1`
		getAllToDos   = `SELECT id, title, created_at
							FROM blogs where 1=1`
	)
	if title != "" {
		getTotalCount = fmt.Sprintf("%s%s", getTotalCount, " and title LIKE '%"+title+"%';")
		getAllToDos = fmt.Sprintf("%s%s", getAllToDos, " and title LIKE '%"+title+"%' ")
	}
	getAllToDos += " ORDER BY created_at OFFSET $1 LIMIT $2;"
	if err := r.db.QueryRowContext(ctx, getTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.QueryRowContext")
	}

	if totalCount == 0 {
		return &models.BlogsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Blogs:      make([]*models.Blog, 0),
		}, nil
	}

	rows, err := r.db.QueryxContext(ctx, getAllToDos, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	blogsList := make([]*models.Blog, 0, query.GetSize())
	for rows.Next() {
		blog := &models.Blog{}
		if err = rows.StructScan(blog); err != nil {
			return nil, errors.Wrap(err, "blogsRepo.GetAll.StructScan")
		}
		blogsList = append(blogsList, blog)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.rows.Err")
	}

	return &models.BlogsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Blogs:      blogsList,
	}, nil
}
