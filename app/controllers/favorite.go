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

// swagger:route GET /story/favorite story getFavorite
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: getFavoriteResponse
//	404: errorResponse
//
// ListAll handles GET requests and returns favorite Story REQUIRED AUTH
func (u *ctrl) HandlerGetFavoriteStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, msg, st, err := u.uc.GetUserFromToken(r)

	ctx, res, msg, st, err := u.uc.GetFavoriteStory(ctx, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

// swagger:route POST /story/favorite story postFavorite
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: postResponse
//	404: errorResponse
//
// ListAll handles POST requests and returns favorite Story REQUIRED AUTH
func (u *ctrl) HandlerPostFavoriteStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var fav models.PostFavoriteStory

	err := json.NewDecoder(r.Body).Decode(&fav)

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.PostFavoriteStory(ctx, &fav, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Favorite error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

// swagger:route DELETE /story/favorite/{storyID} story deleteFavorite
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: postResponse
//	404: errorResponse
//
// ListAll handles DELETE requests and returns favorite Story REQUIRED AUTH
func (u *ctrl) HandlerDeleteFavoriteStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID, _ := strconv.Atoi(chi.URLParam(r, "storyID"))

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.DeleteFavoriteStory(ctx, storyID, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Favorite error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, msg)
}
