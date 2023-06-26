package main

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GeneralResponse struct {
	WithError bool   `json:"with_error"`
	Message   string `json:"message"`
}

type LoginResponse struct {
	GeneralResponse
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// Community estructura de una comunidad
type Community struct {
	// Name nombre de una comunidad. Ej: EDteam
	Name string `json:"name"`
}

// Communities slice de comunidades
type Communities []Community

type Person struct {
	// Name nombre de la persona Ej: Alexys
	Name string `json:"name"`
	// Age edad de la persona Ej: 40
	Age uint8 `json:"age"`
	// Communities comunidades a las que pertenece una persona
	Communities Communities `json:"communities"`
}
