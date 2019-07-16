package api

import "github.com/labstack/echo"

func Run(e *echo.Echo) {

	e.GET("/employee", GETEmployees)
}

func GETEmployees (e echo.Context) error {
	return nil
}