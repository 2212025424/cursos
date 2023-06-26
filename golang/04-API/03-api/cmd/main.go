package main

import (
	"log"
	"net/http"

	"github.com/2212025424/api/authorization"
	"github.com/2212025424/api/handler"
	"github.com/2212025424/api/storage"
)

func main() {

	err := authorization.LoadFiles("C:/Users/enriq/OneDrive/Documentos/educacion/cursos/goland/04-CreacionAPI/03-api/cmd/certificates/app.rsa", "C:/Users/enriq/OneDrive/Documentos/educacion/cursos/goland/04-CreacionAPI/03-api/cmd/certificates/app.rsa.pub")

	if err != nil {
		log.Printf("ERROR: carga de certificados %v\n", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Printf("El servidor se ha ejecutado en el puerto 8080")

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
	}
}
