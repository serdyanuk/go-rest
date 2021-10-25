package core

import (
	"encoding/json"
	"io"
	"net/http"
)

func UnmarshalHTTPBody(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(dest)
	if err == io.EOF {
		return NewAppError(nil, http.StatusBadRequest, "Empty input data")
	}
	return err
}
