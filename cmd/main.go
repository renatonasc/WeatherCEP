package main

import (
	"renatonasc/weathercep/internal/infra/web"
	"renatonasc/weathercep/internal/infra/web/webserver"
)

func main() {

	webserver := webserver.NewWebServer(":8080")

	cepHandler := web.NewCepHandler()

	webserver.AddHandler("GET", "/weather/{cep}", cepHandler.GetWeatherByCep)

	webserver.Start()
}
