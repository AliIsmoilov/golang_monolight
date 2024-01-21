package http

import (
	"github.com/labstack/echo/v4"

	"github.com/AliIsmoilov/golang_monolight/internal/todos"
)

// Map todos routes
func MapToDosRoutes(todoGroup *echo.Group, h todos.Handlers) {
	// docs.SwaggerInfo.Title = cfg.ServiceName
	// docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}
	todoGroup.POST("", h.Create())
	todoGroup.DELETE("/:id", h.Delete())
	todoGroup.PUT("/:id", h.Update())
	todoGroup.GET("/list", h.GetAll())
	todoGroup.GET("/:id", h.GetByID())
}

// Map news routes
func MapNewsRoutes(newsGroup *echo.Group, h todos.NewsHandlers) {
	newsGroup.POST("", h.Create())
	newsGroup.DELETE("/:id", h.Delete())
	newsGroup.DELETE("/soft/:id", h.SoftDelete())
	newsGroup.PUT("/:id", h.Update())
	newsGroup.GET("/list", h.GetAll())
	newsGroup.GET("/:id", h.GetByID())
}
