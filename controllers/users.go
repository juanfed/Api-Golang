package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"v1/models"
	"v1/repositories"
	"v1/service"

	"github.com/labstack/echo/v4"
)

type usercController struct {
	service *service.UserService
}

func NewController(mysql *sql.DB) *usercController {
	return &usercController{
		service: service.NewUserService(
			repositories.NewUserMysqlRepositories(mysql),
			repositories.NewUserRedisRepositories(),
		),
	}
}

func (ctr *usercController) Set(c echo.Context) error {
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctr.service.Set(user))
}

func (ctr *usercController) Get(c echo.Context) error {
	str := c.Param("id")         // digo que ese es el parametro que espero
	id, err := strconv.Atoi(str) // covierto ese parametro a un tipo de dato int
	if err != nil {
		return err
	}

	// mando a las fuciones siguientes el parametro de id
	user, err := ctr.service.Get(id)
	// si hubo un error lo retorno
	if err != nil {
		return err
	}

	// si no hubo error simplemente retorno el usuario
	return c.JSON(http.StatusOK, user)
}

func (ctr *usercController) Update(c echo.Context) error {
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ctr.service.Update(user))
}

func (ctr *usercController) Delete(c echo.Context) error {
	str := c.Param("id")
	// para convertir le dato que me llega como string en int y poder ser enviado a la funcion delete del service
	id, err := strconv.Atoi(str)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ctr.service.Delete(id))
}
