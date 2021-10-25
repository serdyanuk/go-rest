package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/serdyanuk/go-rest/internal/app/core"
)

func UnmarshalHTTPBody(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(dest)
	if err == io.EOF {
		return core.NewAppError(nil, http.StatusBadRequest, "Empty input data")
	}
	return err
}
