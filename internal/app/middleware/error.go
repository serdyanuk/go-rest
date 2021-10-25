package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/serdyanuk/go-rest/internal/app/core"
)

func ErrorHandler(h core.Handler) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := h(rw, r, p)
		if err != nil {
			switch err.(type) {
			case *core.AppError:
			default:

			}
		}
	}
}
