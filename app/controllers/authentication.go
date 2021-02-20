package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Declare a new Person struct.
	var p models.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, res, msg, st, err := u.uc.Registration(ctx, &p)
	if err != nil {
		ctx = logger.Logf(ctx, "user registration error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Declare a new Person struct.
	var p models.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, res, msg, st, err := u.uc.Login(ctx, &p)
	if err != nil {
		ctx = logger.Logf(ctx, "user login error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	header := r.Header.Get("Authorization")
	token := strings.Split(header, " ")

	ctx, res, message, st, err := u.uc.Logout(ctx, token[1])
	if err != nil {
		utils.Response(ctx, w, false, st, message)
		return
	}
	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerCheckSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	header := r.Header.Get("Authorization")
	auth := strings.Split(header, " ")

	user, msg, code, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}

	ctx, res, msg, st, err := u.uc.CheckSession(ctx, user, auth[1])
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerEmailVerification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token := chi.URLParam(r, "token")

	ctx, msg, code, err := u.uc.EmailVerification(ctx, token)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on verify email() => %v", err)
		utils.HTMLResponse(w, code, msg.(string))
		return
	}
	utils.HTMLResponse(w, code, msg.(string))
}
