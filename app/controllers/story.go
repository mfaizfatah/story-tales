package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerInsertStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s models.Story

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, msg, st, err := u.uc.InsertStory(ctx, &s)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

func (u *ctrl) HandlerGetOneStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID := chi.URLParam(r, "storyID")

	ctx, res, msg, st, err := u.uc.GetOneStory(ctx, storyID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetAllStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, res, msg, st, err := u.uc.GetAllStory(ctx)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}
