package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (er AppError) Error() string {
	return er.Message
}

func NewNotFoundError(massage string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: massage,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}
