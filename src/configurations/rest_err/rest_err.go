package rest_err

import (
	"net/http"
)

type Erro struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestError(message, err string, code int, causes []Causes) *Erro {
	return &Erro{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "Bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Erro {
	return &Erro{
		Message: message,
		Err:     "Bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}
func NewInternalError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "Internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}
func NewNotFoundError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "Not_Found",
		Code:    http.StatusNotFound,
	}
}
func NewForbiddenError(message string) *Erro {
	return &Erro{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
func (r *Erro) Error() string {
	return r.Message
}
