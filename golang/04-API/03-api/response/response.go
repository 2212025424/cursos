package response

import (
	"encoding/json"
	"net/http"
)

const (
	MethodNotAllowed           = "Método no permitido"
	DifferentStructureExpected = "La información no tiene la estructura requerida"
	AnIntegerWasExpected       = "Se esperaba un número entero como parámetro"
	SuccessfulProcess          = "Operación completada exitosamente"
	AnErrorHasOccurred         = "Ha ocurrido un error en la operación"
	ResourceNotFound           = "No hay registro con ese ID"
)

type response struct {
	WithError  bool        `json:"with_error"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func New(statusCode int, message string, data interface{}) *response {

	withError := true

	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		withError = false
	}

	return &response{
		withError,
		statusCode,
		message,
		data,
	}
}

func (resp *response) ResponseJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	err := json.NewEncoder(w).Encode(&resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
