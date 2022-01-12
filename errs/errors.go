package errs

import "net/http"

type AppErr struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (r AppErr) AsMessage() *AppErr {
	return &AppErr{Message: r.Message}
}
func NewNotFoundError(msg string) *AppErr {
	return &AppErr{Code: http.StatusNotFound, Message: msg}
}
func NewUnexpectedError(msg string) *AppErr {
	return &AppErr{Code: http.StatusInternalServerError, Message: msg}
}
func NewValidationError(msg string) *AppErr {
	return &AppErr{Code: http.StatusUnprocessableEntity, Message: msg}
}
func NewUnauthorizedError(msg string) *AppErr {
	return &AppErr{Code: http.StatusUnauthorized, Message: msg}
}
