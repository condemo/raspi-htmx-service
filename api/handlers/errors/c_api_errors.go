package errors

import (
	"errors"
	"net/http"
)

type ApiError struct {
	Err    error
	Msg    string
	Status int
}

func NewApiError(err error, msg string, status int) ApiError {
	return ApiError{
		Err:    err,
		Msg:    msg,
		Status: status,
	}
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

var InternalServerError = NewApiError(
	errors.New("internal server error"),
	"internal server errror",
	http.StatusInternalServerError,
)

type UINofifyError struct {
	Msg string
}

func (e UINofifyError) Error() string {
	return e.Msg
}
