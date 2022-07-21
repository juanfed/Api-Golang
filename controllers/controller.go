package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// funcion que me hara la imprecion por consola
func GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, worldx4")
}
