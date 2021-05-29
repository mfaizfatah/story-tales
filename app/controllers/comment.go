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

// swagger:route POST /story/comment story postLikes
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: postResponse
//	404: errorResponse
//
// ListAll handles POST requests and returns coment Story REQUIRED AUTH
func (u *ctrl) HandlerPostComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var comment models.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.PostComment(ctx, &comment, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Comment error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

// swagger:route DELETE /story/comment/{commentID} story deleteComment
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: postResponse
//	404: errorResponse
//
// ListID handles DELETE requests and returns comment Story REQUIRED AUTH
func (u *ctrl) HandlerDeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commentID, _ := strconv.Atoi(chi.URLParam(r, "commentID"))

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.DeleteComment(ctx, commentID, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Comment delete error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, msg)
}

// swagger:route GET /story/favorite story getFavorite
// Return a list of story from the database REQUIRED AUTH
//
// responses:
//	200: getFavoriteResponse
//	404: errorResponse
//
// ListAll handles GET requests and returns favorite Story REQUIRED AUTH
func (u *ctrl) HandlerGetComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	storyID, err := strconv.Atoi(chi.URLParam(r, "storyID"))
	episodeID, err := strconv.Atoi(chi.URLParam(r, "episodeID"))

	ctx, res, msg, st, err := u.uc.GetListComment(ctx, storyID, episodeID)
	if err != nil {
		ctx = logger.Logf(ctx, "List Comment error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetMyComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, res, msg, st, err := u.uc.GetMyComment(ctx, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Get My Comment error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetTopComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	storyID, err := strconv.Atoi(chi.URLParam(r, "storyID"))
	episodeID, err := strconv.Atoi(chi.URLParam(r, "episodeID"))

	ctx, res, msg, st, err := u.uc.GetListTopComment(ctx, storyID, episodeID)
	if err != nil {
		ctx = logger.Logf(ctx, "List TopComment error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerPostLikeComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	commentID, err := strconv.Atoi(chi.URLParam(r, "id"))

	user, msg, code, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}

	ctx, likeCommentSt, msg, st, err := u.uc.GetLikeCommentStatus(ctx, user.ID, commentID)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	if (likeCommentSt != nil) && (likeCommentSt.IDUser == 0) {
		ctx, msg, st, err = u.uc.PostLikeComment(ctx, user.ID, commentID)
		if err != nil {
			ctx = logger.Logf(ctx, "Do LikeComment error() => %v", err)
			utils.Response(ctx, w, false, st, msg)
			return
		}
	} else if likeCommentSt.Deleted == 1 {
		ctx, msg, st, err = u.uc.PostReLikeComment(ctx, user.ID, commentID)
		if err != nil {
			ctx = logger.Logf(ctx, "Do ReLikeComment error() => %v", err)
			utils.Response(ctx, w, false, st, msg)
			return
		}
	} else {
		ctx, msg, st, err = u.uc.PostUnLikeComment(ctx, user.ID, commentID)
		if err != nil {
			ctx = logger.Logf(ctx, "Do UnLikeComment error() => %v", err)
			utils.Response(ctx, w, false, st, msg)
			return
		}
	}

	utils.Response(ctx, w, true, st, msg)
}
