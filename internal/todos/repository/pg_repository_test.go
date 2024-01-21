package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/AliIsmoilov/golang_monolight/internal/models"
)

func TestBlogsRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	commRepo := NewToDosRepository(sqlxDB)

	t.Run("Create", func(t *testing.T) {
		newsUID := uuid.New()
		title := "title"

		rows := sqlmock.NewRows([]string{"id", "title"}).AddRow(newsUID, title)

		blog := &models.Blog{
			ID:        newsUID,
			Title:     title,
			CreatedAt: time.Now(),
		}

		mock.ExpectQuery("").WithArgs(blog.ID, blog.Title).WillReturnRows(rows)

		createdBlog, err := commRepo.Create(context.Background(), blog)

		require.NoError(t, err)
		require.NotNil(t, createdBlog)
		// require.Equal(t, createdBlog, blog)
	})

	t.Run("Create ERR", func(t *testing.T) {
		newsUID := uuid.New()
		title := "title"
		createErr := errors.New("Create blog error")

		blog := &models.Blog{
			ID:    newsUID,
			Title: title,
		}

		mock.ExpectQuery("").WithArgs(blog.ID, blog.Title).WillReturnError(createErr)

		createdToDo, err := commRepo.Create(context.Background(), blog)

		require.Nil(t, createdToDo)
		require.NotNil(t, err)
	})
}

func TestBlogsRepo_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	commRepo := NewToDosRepository(sqlxDB)

	t.Run("Update", func(t *testing.T) {
		blogID := uuid.New()
		title := "title"

		rows := sqlmock.NewRows([]string{"id", "title"}).AddRow(blogID, title)

		blog := &models.Blog{
			ID:    blogID,
			Title: title,
		}

		updatedBlog, err := commRepo.Update(context.Background(), blog)
		mock.ExpectQuery("updateBlog").WithArgs(blog.ID, blog.Title).WillReturnRows(rows)

		require.NoError(t, err)
		require.NotNil(t, updatedBlog)
		require.Equal(t, updatedBlog.ID, blog.ID)
	})

	t.Run("Update ERR", func(t *testing.T) {
		blogID := uuid.New()
		title := "title"

		blog := &models.Blog{
			ID:    blogID,
			Title: title,
		}

		updatedBlog, err := commRepo.Update(context.Background(), blog)

		require.NotNil(t, err)
		require.Nil(t, updatedBlog)
	})
}

func TestBlogsRepo_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	commRepo := NewToDosRepository(sqlxDB)

	t.Run("Delete", func(t *testing.T) {
		blogID := uuid.New()
		mock.ExpectExec("deleteBlog").WithArgs(blogID).WillReturnResult(sqlmock.NewResult(1, 1))
		err := commRepo.Delete(context.Background(), blogID)

		require.NoError(t, err)
	})

	t.Run("Delete Err", func(t *testing.T) {
		blogID := uuid.New()

		mock.ExpectExec("deleteBlog").WithArgs(blogID).WillReturnResult(sqlmock.NewResult(1, 0))

		err := commRepo.Delete(context.Background(), blogID)
		require.NotNil(t, err)
	})
}
