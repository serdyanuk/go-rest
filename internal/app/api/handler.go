package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error
