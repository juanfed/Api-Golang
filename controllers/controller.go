package controllers

import (
	"fmt"
	"net/http"

	"v1/models"

	"github.com/labstack/echo/v4"
)

// funcion que me hara la imprecion por consola
func GetHello(c echo.Context) error {
	// en caso de que no exista un error y la opracion halla sido correcta mando un hola mundo
	return c.String(http.StatusOK, "Hello, worldx5")
}

func GetUser(c echo.Context) error {
	// creo un nuevo usuario en base al modelo de que espero recibir
	user := models.User{}

	// instancio el error que me llega de Bind(&user), Bind en este caso me sirve para decir que debo de recibir una solicitud con los datos espesificados en el modelo de user
	err := c.Bind(&user)

	// imprimo el usuario en consola
	fmt.Println(user)

	// compruebo de que no exista un error, en caso de que si lo capturo y lo muestro al usuario
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// mostrare los usuarios, en caso de que no exista un error en la peticion
	return c.JSON(http.StatusOK, user)
}
