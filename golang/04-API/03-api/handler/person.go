package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/2212025424/api/model"
	"github.com/2212025424/api/response"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.New(http.StatusBadRequest, response.MethodNotAllowed, nil).ResponseJSON(w)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.New(http.StatusBadRequest, response.DifferentStructureExpected, nil).ResponseJSON(w)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response.New(http.StatusInternalServerError, response.AnErrorHasOccurred, nil).ResponseJSON(w)
		return
	}

	response.New(http.StatusCreated, response.SuccessfulProcess, nil).ResponseJSON(w)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response.New(http.StatusBadRequest, response.MethodNotAllowed, nil).ResponseJSON(w)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.New(http.StatusBadRequest, response.AnIntegerWasExpected, nil).ResponseJSON(w)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.New(http.StatusBadRequest, response.AnIntegerWasExpected, nil).ResponseJSON(w)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response.New(http.StatusInternalServerError, response.AnErrorHasOccurred, nil).ResponseJSON(w)
		return
	}

	response.New(http.StatusOK, response.SuccessfulProcess, nil).ResponseJSON(w)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.New(http.StatusBadRequest, response.SuccessfulProcess, nil).ResponseJSON(w)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.New(http.StatusBadRequest, response.AnIntegerWasExpected, nil).ResponseJSON(w)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response.New(http.StatusBadRequest, response.ResourceNotFound, nil).ResponseJSON(w)
		return
	}
	if err != nil {
		response.New(http.StatusInternalServerError, response.AnErrorHasOccurred, nil).ResponseJSON(w)
		return
	}

	response.New(http.StatusOK, response.SuccessfulProcess, nil).ResponseJSON(w)
}

func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.New(http.StatusBadRequest, response.MethodNotAllowed, nil).ResponseJSON(w)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.New(http.StatusBadRequest, response.AnIntegerWasExpected, nil).ResponseJSON(w)
		return
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response.New(http.StatusBadRequest, response.ResourceNotFound, nil).ResponseJSON(w)
		return
	}
	if err != nil {
		response.New(http.StatusInternalServerError, response.AnErrorHasOccurred, nil).ResponseJSON(w)
		return
	}

	response.New(http.StatusOK, response.SuccessfulProcess, data).ResponseJSON(w)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.New(http.StatusBadRequest, response.MethodNotAllowed, nil).ResponseJSON(w)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response.New(http.StatusInternalServerError, response.AnErrorHasOccurred, nil).ResponseJSON(w)
		return
	}

	response.New(http.StatusOK, response.SuccessfulProcess, &data).ResponseJSON(w)
}
