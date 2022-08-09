package handler

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
)

type (
	authRequest struct {
		Email    string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	authResponse struct {
		Token string `json:"token"`
	}
	registerRequest struct {
		Username  string                `form:"username" binding:"required"`
		Password  string                `form:"password" binding:"required"`
		Photo     *multipart.FileHeader `form:"photo"`
		FirstName string                `form:"first_name"`
		LastName  string                `form:"last_name"`
		Email     string                `form:"email"`
	}
)

var (
	errNoSuchUser        = errors.New("no such user")
	errUserAlreadyExists = errors.New("user already exists")
)

func (s *Server) auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request authRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	user, err := s.userStore.GetByEmail(ctx, request.Email)
	if err != nil {
		s.respond.Error(w, http.StatusBadRequest, err)
		return
	}

}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
}
