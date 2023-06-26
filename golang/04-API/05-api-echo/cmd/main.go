package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/2212025424/api/authorization"
	"github.com/2212025424/api/handler"
	"github.com/2212025424/api/storage"
)

func main() {
	err := authorization.LoadFiles("C:/Users/enriq/OneDrive/Documentos/educacion/cursos/goland/04-API/05-api-echo/cmd/certificates/app.rsa", "C:/Users/enriq/OneDrive/Documentos/educacion/cursos/goland/04-API/05-api-echo/cmd/certificates/app.rsa.pub")

	if err != nil {
		log.Printf("ERROR: carga de certificados %v\n", err)
	}

	store := storage.NewMemory()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	log.Println("Servidor iniciado en el puerto 8080")

	err = e.Start(":8080")

	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}

}
