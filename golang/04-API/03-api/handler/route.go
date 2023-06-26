package handler

import (
	"net/http"

	"github.com/2212025424/api/middleware"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", middleware.Authentication(h.create))
	mux.HandleFunc("/v1/persons/getall", h.getAll)
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", h.delete)
	mux.HandleFunc("/v1/persons/gtbyid", h.getByID)
}

func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/v1/login", h.login)
}
