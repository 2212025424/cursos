package response

const (
	MethodNotAllowed           = "Método no permitido"
	DifferentStructureExpected = "La información no tiene la estructura requerida"
	AnIntegerWasExpected       = "Se esperaba un número entero como parámetro"
	AnErrorHasOccurred         = "Ha ocurrido un error en la operación"
	ResourceNotFound           = "No hay registro con ese ID"
	IncorrectCredentials       = "Las credenciales no son correctas"
	TokenNotCreated            = "Error al genrar el token"
	IdentifierNotExist         = "No hay registro con este identificador"
	SuccessfulProcess          = "Operación completada exitosamente"
)

type response struct {
	WithError bool        `json:"with_error"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func New(message string, data interface{}) response {

	withError := false

	if message != SuccessfulProcess {
		withError = true
	}

	return response{
		withError,
		message,
		data,
	}
}
