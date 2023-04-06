package user

import (
	"fmt"
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/RustamRR/job-rest-api/internal/store"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

var serverStore store.Store

func InitRoutes(e *echo.Echo, store store.Store) {
	serverStore = store

	e.POST("/users", create)
	e.GET("/users/:id", get)
	e.GET("/users", getAll)
	e.PATCH("/users/:id", update)
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

	return c.JSON(http.StatusCreated, user)
}

func get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "не передан id")
	}

	userUuid, errParseUuid := uuid.Parse(id)
	if errParseUuid != nil {
		return c.JSON(http.StatusBadRequest, "некорректный id")
	}

	user, err := serverStore.User().Find(userUuid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func getAll(c echo.Context) error {
	users, err := serverStore.User().FindAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func update(c echo.Context) error {
	var userUpdate model.UserUpdate
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "не передан id")
	}

	userUuid, errParseUuid := uuid.Parse(id)
	if errParseUuid != nil {
		return c.JSON(http.StatusBadRequest, "некорректный id")
	}

	user, err := serverStore.User().Find(userUuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "пользователь не найден")
	}

	if err := c.Bind(&userUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("некорректные данные"))
	}

	model.UpdateUser(&user, &userUpdate)

	if err := serverStore.User().Update(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("ошибка обновления данных: %v", err))
	}

	return c.JSON(http.StatusOK, user)
}
