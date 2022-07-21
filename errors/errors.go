package errors

import "net/http"

type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	return e.Message
}

func NewInvalidRequest(msg string) *Error {
	return New(http.StatusBadRequest, msg)
}

func NewInternalServer(msg string) *Error {
	return New(http.StatusInternalServerError, msg)
}

func NewExpected(msg string) *Error {
	return New(http.StatusOK, msg)
}

func New(statusCode int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    msg,
	}
}
