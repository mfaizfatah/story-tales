package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerPostStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s models.Story

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.PostStory(ctx, &s, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

func (u *ctrl) HandlerGetOneStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID, err := strconv.Atoi(chi.URLParam(r, "storyID"))

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

func (u *ctrl) HandlerGetDetailEpisode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID, err := strconv.Atoi(chi.URLParam(r, "storyID"))
	episodeID, err := strconv.Atoi(chi.URLParam(r, "episodeID"))
	ctx, res, msg, st, err := u.uc.GetDetailEpisode(ctx, storyID, episodeID)
	if err != nil {
		ctx = logger.Logf(ctx, "Detail error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}
