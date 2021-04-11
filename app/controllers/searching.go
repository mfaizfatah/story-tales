package controllers

import (
	"net/http"
	"time"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerSearching(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query().Get("searchQuery")
	genre := r.URL.Query().Get("genre")
	ctx, res, msg, st, err := u.uc.Searching(ctx, query, genre)
	if err != nil || st >= 400 {
		ctx = logger.Logf(ctx, "error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGenerateDoc(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	for i := 1; i <= 3; i++ {
		ctx, res, msg, st, err := u.uc.GenerateDocument(ctx)
		if err != nil || st >= 400 {
			time.Sleep(time.Duration(i) * time.Minute)
		} else if i == 3 && st >= 400 {
			ctx = logger.Logf(ctx, "error while generate search document() => %v", err)
			utils.Response(ctx, w, false, st, msg)
			return
		} else {
			utils.Response(ctx, w, true, st, res)
			return
		}
	}
}
