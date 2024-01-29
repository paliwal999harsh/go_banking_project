package errs

import "net/http"

type AppError struct {
	Code    int
	Message string `json:"message"`
}

func (e AppError) AsMessage() interface{} {
	return e.Message
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		http.StatusNotFound,
		message,
	}
}
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		http.StatusInternalServerError,
		message,
	}
}
func NewBadRequestError(message string) *AppError {
	return &AppError{
		http.StatusUnprocessableEntity,
		message,
	}
}
