package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// aca voy a iniciar las rutas del servidor
func StartRoutes() {
	// empizo por iniciar las rutas

	//decalro la ruta y lo que se ejecutara cuando se ingrese a ella, siempre declaro las rutas antes de levantar el servidor
	e.GET("/", GetHello)
}

// funcion que me hara la imprecion por consola
func GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, worldx3")
}
