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

func (u *ctrl) HandlerPostLikes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var likes models.Likes

	err := json.NewDecoder(r.Body).Decode(&likes)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.PostLikes(ctx, &likes, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Likes error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

func (u *ctrl) HandlerDeleteLikes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID, _ := strconv.Atoi(chi.URLParam(r, "storyID"))
	episodeID, _ := strconv.Atoi(chi.URLParam(r, "episodeID"))

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.DeleteLikes(ctx, storyID, episodeID, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Likes error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, msg)
}
