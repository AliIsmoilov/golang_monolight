package todos

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type Handlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

// News HTTP Handlers interface
type NewsHandlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	SoftDelete() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}
