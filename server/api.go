package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//instancio un nuevo objeto
var e = echo.New()

// donde inicio el servidor
func Start() {
	// empizo por iniciar las rutas

	//decalro la ruta y lo que se ejecutara cuando se ingrese a ella, siempre declaro las rutas antes de levantar el servidor
	e.GET("/", GetHello)

	// Digo que levante el servirdor en el puerto 3000
	e.Logger.Fatal(e.Start(":3000"))
}

// funcion que me hara la imprecion por consola
func GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, worldx2")
}
