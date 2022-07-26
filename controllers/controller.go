package controllers

import (
	"net/http"
	"strconv"
	"v1/models"
	"v1/repositories"
	"v1/service"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service *service.UserService
}

func NewController() *controller {
	return &controller{
		service: service.NewUserService(
			repositories.NewUserMysqlRepositories(),
			repositories.NewUserRedisRepositories(),
		),
	}
}
func (ctr *controller) SetUser(c echo.Context) error {
	user := models.User{}

	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctr.service.SetUser(user))
}
func (ctr *controller) DeleteUser(c echo.Context) error {
	str := c.Param("id")
	// para convertir le dato que me llega como string en int y poder ser enviado a la funcion delete del service
	id, err := strconv.Atoi(str)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ctr.service.DeleteUser(id))
}
