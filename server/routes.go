package server

import (
	"v1/controllers"
)

// aca voy a iniciar las rutas del servidor
func StartRoutes() {
	// empizo por iniciar las rutas

	//decalro la ruta y lo que se ejecutara cuando se ingrese a ella, siempre declaro las rutas antes de levantar el servidor
	e.GET("/", controllers.GetHello)
}
