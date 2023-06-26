package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/2212025424/api/model"
	"github.com/2212025424/api/response"
	"github.com/labstack/echo"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {

	data := model.Person{}

	err := c.Bind(&data)

	if err != nil {
		resp := response.New(response.DifferentStructureExpected, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = p.storage.Create(&data)
	if err != nil {
		resp := response.New(response.AnErrorHasOccurred, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := response.New(response.SuccessfulProcess, nil)
	return c.JSON(http.StatusOK, resp)
}

func (p *person) update(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := response.New(response.AnIntegerWasExpected, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	data := model.Person{}

	err = c.Bind(&data)

	if err != nil {
		resp := response.New(response.DifferentStructureExpected, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = p.storage.Update(ID, &data)

	if err != nil {
		resp := response.New(response.AnErrorHasOccurred, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := response.New(response.SuccessfulProcess, nil)
	return c.JSON(http.StatusOK, resp)
}

func (p *person) delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := response.New(response.AnIntegerWasExpected, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = p.storage.Delete(ID)

	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		resp := response.New(response.IdentifierNotExist, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		resp := response.New(response.AnErrorHasOccurred, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := response.New(response.SuccessfulProcess, nil)
	return c.JSON(http.StatusInternalServerError, resp)
}

func (p *person) getByID(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := response.New(response.AnIntegerWasExpected, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, err := p.storage.GetByID(ID)

	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		resp := response.New(response.IdentifierNotExist, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err != nil {
		resp := response.New(response.AnErrorHasOccurred, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := response.New(response.SuccessfulProcess, data)
	return c.JSON(http.StatusOK, resp)
}

func (p *person) getAll(c echo.Context) error {

	data, err := p.storage.GetAll()

	if err != nil {
		resp := response.New(response.AnErrorHasOccurred, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := response.New(response.SuccessfulProcess, &data)
	return c.JSON(http.StatusOK, resp)
}
