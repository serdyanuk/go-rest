package api

import (
	"encoding/json"
	"net/http"
)

type APIErrorMessage struct {
	Error APIErrorMessageContent `json:"error"`
}

type APIErrorMessageContent struct {
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, code int, message string) error {
	return SendJSON(w, code, APIErrorMessage{
		APIErrorMessageContent{
			Message: message,
		},
	})
}

func SendJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.WriteHeader(code)
	w.Header().Add("Content-type", "application/json")

	return json.NewEncoder(w).Encode(v)
}
