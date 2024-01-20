package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/AliIsmoilov/golang_monolight/internal/models"
	"github.com/AliIsmoilov/golang_monolight/pkg/httpErrors"
	"github.com/AliIsmoilov/golang_monolight/pkg/utils"
)

// Create
// @Summary Create new news
// @Description CreateNews new news
// @Tags NewsHeader
// @Accept  json
// @Produce  json
// @Param body body models.NewsSwagger true "body"
// @Success 201 {object} models.News
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/news [post]
func (h *blogHandlers) CreateNews() echo.HandlerFunc {
	return func(c echo.Context) error {

		news := &models.News{}

		if err := utils.SanitizeRequest(c, news); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		createdBlog, err := h.todosUC.CreateNews(c.Request().Context(), news)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdBlog)
	}
}
