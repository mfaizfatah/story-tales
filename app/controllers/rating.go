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

// swagger:route POST /story/rating story postRating
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: postResponse
//	404: errorResponse
//
// ListAll handles POST requests and returns rating Story REQUIRED AUTH
func (u *ctrl) HandlerPostRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s models.Rating

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

	ctx, msg, st, err = u.uc.PostRating(ctx, &s, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Rating error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

// swagger:route DELETE /story/rating/{storyID} story deleteRating
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: postResponse
//	404: errorResponse
//
// ListID handles DELETE requests and returns RATING Story REQUIRED AUTH
func (u *ctrl) HandlerDeleteRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID, _ := strconv.Atoi(chi.URLParam(r, "storyID"))

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.DeleteRating(ctx, storyID, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Likes error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, msg)
}
