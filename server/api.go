package server

import (
	"github.com/labstack/echo/v4"
)

//instancio un nuevo objeto
var e = echo.New()

// donde inicio el servidor
func Start() {
	// levanto primero las rutas por donde se daran las peticiones del servidor
	StartRoutes()
	// Digo que levante el servirdor en el puerto 3000
	e.Logger.Fatal(e.Start(":3000"))
}
