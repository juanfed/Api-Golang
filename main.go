// creo el archivo main, sera el encargado de encencer mi servio, ya que es el archivo principal que se irá a ejecutar
package main

import (
	"v1/server"
)

func main() {
	// llamo a la funcion que me iniciara el servidor en la ruta que haya designado
	server.Start()
}
