package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover()) // El servidor trata de recuperarse de una excepcion

	e.GET("/", saludar)
	e.GET("/dividir", dividir)

	persona := e.Group("/persona")
	persona.POST("", crear)
	persona.GET("/:id", consultar)
	persona.PUT("/:id", actualizar)
	persona.DELETE("/:id", eliminar)

	e.Start(":8080")
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "que onda !!"})
}

func dividir(c echo.Context) error {
	f := c.QueryParam("id")

	i, _ := strconv.Atoi(f)

	r := 100 / i

	return c.String(http.StatusOK, "res: "+strconv.Itoa(r))
}

func crear(c echo.Context) error {
	return c.String(http.StatusOK, "creado")
}

func eliminar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "eliminar: "+id)
}

func actualizar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "actualizar: "+id)
}

func consultar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "consultar: "+id)
}
