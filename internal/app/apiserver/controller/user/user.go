package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRoutes(e *echo.Echo) {
	g := e.Group("/users")

	g.POST("/create", create)
}

func create(c echo.Context) error {
	return c.String(http.StatusCreated, "ok")
}
