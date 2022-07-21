// creo el archivo main, sera el encargado de encencer mi servio, ya que es el archivo principal que se irá a ejecutar
package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//instancio un nuevo objeto
var e = echo.New()

//funcion que cre inicia el servirdor
func Start() {
	//decalro la ruta y lo que se ejecutara cuando se ingrese a ella, siempre declaro las rutas antes de levantar el servidor
	e.GET("/", getHello)

	// Digo que levante el servirdor en el puerto 3000
	e.Logger.Fatal(e.Start(":3000"))

}

func main() {
	fmt.Println("Hello from main ")
	// llamo a la funcion que me iniciara el servidor en la ruta que haya designado
	Start()
}

func getHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world")
}
