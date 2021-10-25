package core

import "net/http"

type AppError struct {
	err      error
	httpCode int
	message  string
}

func NewAppError(err error, httpCode int, message string) error {
	return &AppError{err, httpCode, message}
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) Unwrap() error {
	return e.err
}

func NewSystemError(err error, message string) error {
	return NewAppError(err, http.StatusInternalServerError, message)
}
