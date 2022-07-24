package server

import "github.com/labstack/echo/v4"

var e = echo.New()

func Start() {
	StartRoutes()
	e.Logger.Fatal(e.Start(":3000"))
}
