package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/serdyanuk/go-rest/internal/app/api"
)

func ErrorHandler(h api.Handler) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := h(rw, r, p)
		if err != nil {

		}
	}
}
