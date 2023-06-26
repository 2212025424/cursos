package handler

import (
	"net/http"

	"github.com/labstack/echo"

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

func (l *login) login(c echo.Context) error {

	data := model.Login{}

	err := c.Bind(&data)

	if err != nil {
		resp := response.New(response.DifferentStructureExpected, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !isLoginValid(&data) {
		resp := response.New(response.IncorrectCredentials, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := response.New(response.TokenNotCreated, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}

	resp := response.New(response.SuccessfulProcess, dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@ed.team" && data.Password == "123456"
}
