package main

import (
	"github.com/jhonata-menezes/luizalabs-desafio/back/api"
	"github.com/jhonata-menezes/luizalabs-desafio/back/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8001"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodDelete},
	}))

	api.LoadRoutes(e)
	log.Println("listening ", config.GetConfig().App.Host)
	http.ListenAndServe(config.GetConfig().App.Host, e)
}
