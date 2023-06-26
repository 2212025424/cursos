package handler

import (
	"encoding/json"
	"net/http"

	"github.com/2212025424/api/authorization"
	"github.com/2212025424/api/model"
	"github.com/2212025424/api/response"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.New(http.StatusBadRequest, response.MethodNotAllowed, nil).ResponseJSON(w)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.New(http.StatusBadRequest, response.DifferentStructureExpected, nil).ResponseJSON(w)
		return
	}

	if !isLoginValid(&data) {
		response.New(http.StatusBadRequest, "contraseña no válida", nil).ResponseJSON(w)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response.New(http.StatusInternalServerError, "No se pudo generar el token", nil).ResponseJSON(w)
		return
	}

	dataToken := map[string]string{"token": token}

	response.New(http.StatusOK, response.SuccessfulProcess, dataToken).ResponseJSON(w)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ed.team" && data.Password == "123456"
}
