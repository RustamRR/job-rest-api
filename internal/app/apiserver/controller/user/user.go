package user

import (
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/RustamRR/job-rest-api/internal/store"
	"github.com/labstack/echo/v4"
	"net/http"
)

var serverStore store.Store

func InitRoutes(e *echo.Echo, store store.Store) {
	serverStore = store
	g := e.Group("/users")

	g.POST("/create", create)
}

func create(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := serverStore.User().Create(user)
	if err != nil {
		return err
	}
	user.Sanitize()

	return c.JSON(http.StatusCreated, user)
}
