package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/2212025424/api/response"
	"github.com/labstack/echo"
)

func TestPerson_Create_wrong_structure(t *testing.T) {
	data := []byte(`{"name": 123, "age":18}`)

	w := httptest.NewRecorder() // return a writer
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockOk{})

	err := p.create(ctx)

	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo: %v", err)
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("Código estado: se esperaba: %d, se obtuvo: %d", http.StatusBadRequest, w.Code)
	}

	type temstruct struct {
		WithError bool        `json:"with_error"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
	}

	resp := temstruct{}
	err = json.NewDecoder(w.Body).Decode(&resp)

	if err != nil {
		t.Errorf("Error unmarshal Body: %v", err)
	}

	if resp.Message != response.DifferentStructureExpected {
		t.Errorf("Respuesta no esperada: %q - Se esperaba: %q", resp.Message, response.DifferentStructureExpected)
	}

}

func TestPerson_Create_wrong_storage(t *testing.T) {
	data := []byte(`{"name": "123", "age":18}`)

	w := httptest.NewRecorder() // return a writer
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockError{})

	err := p.create(ctx)

	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo: %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Código estado: se esperaba: %d, se obtuvo: %d", http.StatusInternalServerError, w.Code)
	}

	type temstruct struct {
		WithError bool        `json:"with_error"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
	}

	resp := temstruct{}
	err = json.NewDecoder(w.Body).Decode(&resp)

	if err != nil {
		t.Errorf("Error unmarshal Body: %v", err)
	}

	if resp.Message != response.AnErrorHasOccurred {
		t.Errorf("Respuesta no esperada: %q - Se esperaba: %q", resp.Message, response.AnErrorHasOccurred)
	}
}

func TestPerson_Create(t *testing.T) {
	data := []byte(`{"name": "123", "age":18}`)

	w := httptest.NewRecorder() // return a writer
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-type", "application/json")

	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&StorageMockOk{})

	err := p.create(ctx)

	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo: %v", err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Código estado: se esperaba: %d, se obtuvo: %d", http.StatusCreated, w.Code)
	}

	type temstruct struct {
		WithError bool        `json:"with_error"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
	}

	resp := temstruct{}
	err = json.NewDecoder(w.Body).Decode(&resp)

	if err != nil {
		t.Errorf("Error unmarshal Body: %v", err)
	}

	if resp.Message != response.SuccessfulProcess {
		t.Errorf("Respuesta no esperada: %q - Se esperaba: %q", resp.Message, response.SuccessfulProcess)
	}
}
