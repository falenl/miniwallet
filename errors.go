package error

import "net/http"

type errorServer struct {
	StatusCode int
	Message    string
}

func (e *errorServer) Error() string {
	return e.Message
}

func NewErrInvalidRequest(msg string) errorServer {
	return errorServer{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}
