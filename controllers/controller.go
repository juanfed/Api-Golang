package controllers

import (
	"fmt"
	"net/http"
	"v1/models"
	"v1/repositories"
	"v1/service"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service      *service.UserService
	serviceRedis *service.UserService
}

func NewController() *controller {
	return &controller{
		service: service.NewUserService(
			repositories.NewUserPostgresRepositories(),
		),
		serviceRedis: service.NewUserServiceRedis(
			repositories.NewUserRedisRepositories(),
		),
	}
}
func (ctr *controller) SetUser(c echo.Context) error {
	user := models.User{}

	err := c.Bind(&user)
	fmt.Println("Datos de user: ", user)
	fmt.Println("Id del user: ", user.Id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ctr.serviceRedis.SetUserRedis(user)
	return c.JSON(http.StatusOK, ctr.service.SetUser(user))
}
