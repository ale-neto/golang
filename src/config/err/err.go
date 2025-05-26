package err_rest

import "net/http"

type Err struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *Err) Error() string {
	return e.Message
}

func NewErr(message string, err string, code int, causes []Causes) *Err {
	return &Err{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *Err {
	return &Err{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationErr(message string, causes []Causes) *Err {
	return &Err{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerErr(message string) *Err {
	return &Err{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundErr(message string) *Err {
	return &Err{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func NewUnauthorizedErr(message string) *Err {
	return &Err{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewForbiddenErr(message string) *Err {
	return &Err{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewInternalServerError(message string) *Err {
	return &Err{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}
