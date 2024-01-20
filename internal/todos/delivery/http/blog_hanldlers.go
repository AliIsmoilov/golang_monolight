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

// Blog handlers
type blogHandlers struct {
	cfg     *config.Config
	todosUC todos.UseCase
	logger  logger.Logger
}

// NewBlogHandlers Blog handlers constructor
func NewBlogHandlers(cfg *config.Config, todosUC todos.UseCase, logger logger.Logger) todos.Handlers {
	return &blogHandlers{cfg: cfg, todosUC: todosUC, logger: logger}
}

// CreateBlog
// @Summary CreateBlog new blog
// @Description create new blog
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param body body models.BlogSwagger true "body"
// @Success 201 {object} models.Blog
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs [post]
func (h *blogHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		blog := &models.Blog{}

		if err := utils.SanitizeRequest(c, blog); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		createdBlog, err := h.todosUC.Create(c.Request().Context(), blog)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdBlog)
	}
}

// Update
// @Summary Update blog
// @Description update new blog
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param body body models.BlogSwagger true "body"
// @Success 200 {object} models.BlogSwagger
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/{id} [put]
func (h *blogHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		comm := &models.Blog{}
		if err = utils.SanitizeRequest(c, comm); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updatedToDo, err := h.todosUC.Update(c.Request().Context(), &models.Blog{
			ID:    blogsID,
			Title: comm.Title,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedToDo)
	}
}

// Delete
// @Summary Delete blog
// @Description delete blog
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/{id} [delete]
func (h *blogHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.todosUC.Delete(c.Request().Context(), blogsID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Get blog
// @Description Get blog by id
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Blog
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/{id} [get]
func (h *blogHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		blog, err := h.todosUC.GetByID(c.Request().Context(), blogsID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, blog)
	}
}

// GetAll
// @Summary Get Blog
// @Description Get all blog
// @Tags Blog
// @Accept  json
// @Produce  json
// @Param title query string false "title"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Success 200 {object} models.BlogsList
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/list [get]
func (h *blogHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		toDoList, err := h.todosUC.GetAll(c.Request().Context(), c.QueryParam("title"), pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, toDoList)
	}
}
