package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/serdyanuk/go-rest/internal/app/core"
	"github.com/serdyanuk/go-rest/internal/app/store"
)

func Signup(userRepo store.UserRepostitory) core.Handler {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
		signupInput := SignupInput{}
		err := core.UnmarshalHTTPBody(r, &signupInput)
		if err != nil {
			return err
		}
		user, err := userRepo.Create(signupInput.Login, signupInput.Password)
		if err != nil {
			return err
		}
		_ = user
		return nil
	}
}

type SignupInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}