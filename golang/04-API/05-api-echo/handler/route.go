package handler

import (
	"github.com/2212025424/api/middleware"
	"github.com/labstack/echo"
)

func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)

	person := e.Group("/v1/person")

	person.Use(middleware.Authentication)

	person.POST("", h.create)
	person.PUT("/:id", h.update)
	person.DELETE("/:id", h.delete)
	person.GET("/:id", h.getByID)
	person.GET("", h.getAll)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
