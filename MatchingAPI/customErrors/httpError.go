package customErrors

import "net/http"

type HttpError struct {
	StatusCode  int
	ErrorString string
}

func (e *HttpError) Error() string {
	return e.ErrorString
}

func (e *HttpError) SendResponse(w http.ResponseWriter) {
	w.WriteHeader(e.StatusCode)
	w.Write([]byte(e.ErrorString))
}
