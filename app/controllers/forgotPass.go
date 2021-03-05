package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerSendEmailForgotPass(w http.ResponseWriter, r *http.Request) {
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

	ctx, res, msg, code, err := u.uc.SendLinkForgotPass(ctx, &p)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on send link forgot password() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}
	utils.Response(ctx, w, true, code, res)
}

func (u *ctrl) HandlerValidateTokenForgotPass(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token := chi.URLParam(r, "token")

	ctx, msg, code, err := u.uc.ValidateTokenForgotPass(ctx, token)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on validate token forgot password() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}
	utils.Response(ctx, w, true, code, msg)
}

func (u *ctrl) HandlerChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Declare a new Forgot Pass struct.
	var p models.ForgotPass

	// get user id from token
	user, msg, code, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, msg, st, err := u.uc.ChangePassword(ctx, user.ID, &p)
	if err != nil {
		ctx = logger.Logf(ctx, "error while change password() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, msg)
}
