package validations

import "errors"

var (
	// ErrorRowsAffected .
	ErrorRowsAffected = errors.New("Se afecta más de una fila")
)
