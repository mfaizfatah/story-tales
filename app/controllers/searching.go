package controllers

import (
	"net/http"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerSearching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query().Get("searchQuery")
	ctx, res, msg, st, err := u.uc.Searching(ctx, query)
	if err != nil || st >= 400 {
		ctx = logger.Logf(ctx, "error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}
