package server

import (
	"net/http"
	"strings"

	"github.com/AliIsmoilov/golang_monolight/docs"
	"github.com/AliIsmoilov/golang_monolight/pkg/csrf"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	todosHttp "github.com/AliIsmoilov/golang_monolight/internal/todos/delivery/http"
	todosRepository "github.com/AliIsmoilov/golang_monolight/internal/todos/repository"
	todosUseCase "github.com/AliIsmoilov/golang_monolight/internal/todos/usecase"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:8080
// @BasePath /api/v1
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	cRepo := todosRepository.NewToDosRepository(s.db)
	commUC := todosUseCase.NewToDosUseCase(s.cfg, cRepo, s.logger)

	// Init handlers
	todoHandlers := todosHttp.NewBlogHandlers(s.cfg, commUC, s.logger)
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "ToDo API"
	docs.SwaggerInfo.Description = "ToDo REST API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.Start(":5050")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	v1 := e.Group("/v1")

	health := v1.Group("/health")
	todoGroup := v1.Group("/blogs")

	todosHttp.MapToDosRoutes(todoGroup, todoHandlers)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy!"})
	})

	return nil
}
