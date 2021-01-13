package controllers

import (
	"net/http"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerSendEmailForgotPass(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, msg, code, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}

	ctx, res, msg, code, err := u.uc.SendLinkForgotPass(ctx, user)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}
	utils.Response(ctx, w, true, code, res)
}
