package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/AliIsmoilov/golang_monolight/config"
	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/AliIsmoilov/golang_monolight/internal/todos"
	"github.com/AliIsmoilov/golang_monolight/pkg/httpErrors"
	"github.com/AliIsmoilov/golang_monolight/pkg/logger"
	"github.com/AliIsmoilov/golang_monolight/pkg/utils"
)

// News handlers
type newsHandlers struct {
	cfg    *config.Config
	newsUC todos.NewsUseCase
	logger logger.Logger
}

// NewBlogHandlers Blog handlers constructor
func NewNewsHandlers(cfg *config.Config, newsUC todos.NewsUseCase, logger logger.Logger) todos.NewsHandlers {
	return &newsHandlers{cfg: cfg, newsUC: newsUC, logger: logger}
}

// Create
// @Summary Create new news
// @Description CreateNews new news
// @Tags News
// @Accept  json
// @Produce  json
// @Param body body models.NewsSwagger true "body"
// @Success 201 {object} models.News
// @Failure 500 {object} httpErrors.RestErr
// @Router /news [post]
func (h *newsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		news := &models.News{}

		if err := utils.SanitizeRequest(c, news); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		createdBlog, err := h.newsUC.Create(c.Request().Context(), news)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdBlog)
	}
}

// Update
// @Summary Update news
// @Description update news
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param body body models.NewsSwagger true "body"
// @Success 200 {object} models.NewsSwagger
// @Failure 500 {object} httpErrors.RestErr
// @Router /news/{id} [put]
func (h *newsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		newsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		comm := &models.News{}
		if err = utils.SanitizeRequest(c, comm); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updatedNews, err := h.newsUC.Update(c.Request().Context(), &models.News{
			ID:          newsID,
			Title:       comm.Title,
			Description: comm.Description,
			Photo:       comm.Photo,
			PublishedBy: comm.PublishedBy,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedNews)
	}
}

// Delete
// @Summary Delete news
// @Description delete news
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Router /news/{id} [delete]
func (h *newsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		newsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.newsUC.Delete(c.Request().Context(), newsID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Get news
// @Description Get news by id
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.News
// @Failure 500 {object} httpErrors.RestErr
// @Router /news/{id} [get]
func (h *newsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		newsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		news, err := h.newsUC.GetByID(c.Request().Context(), newsID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, news)
	}
}

// GetAll
// @Summary Get News
// @Description Get all news
// @Tags News
// @Accept  json
// @Produce  json
// @Param title query string false "title"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Success 200 {object} models.NewsList
// @Failure 500 {object} httpErrors.RestErr
// @Router /news/list [get]
func (h *newsHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		newsList, err := h.newsUC.GetAll(c.Request().Context(), c.QueryParam("title"), pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, newsList)
	}
}

// Soft Delete
// @Summary Soft Delete news
// @Description soft delete news
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Router /news/soft/{id} [delete]
func (h *newsHandlers) SoftDelete() echo.HandlerFunc {
	return func(c echo.Context) error {

		newsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.newsUC.SoftDelete(c.Request().Context(), newsID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}
