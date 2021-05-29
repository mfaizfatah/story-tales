package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerGetFollow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	user, msg, code, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}
	ctx, followSt, msg, st, err := u.uc.GetFollowStatus(ctx, user.ID, id)
	if err != nil {
		ctx = logger.Logf(ctx, "Get Follow Status error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	var stFoll = new(models.StatusFollow)

	if ((followSt != nil) && (followSt.UserFollowID == 0)) || followSt.Deleted == 1 {
		stFoll.StatusFollow = false
	} else {
		stFoll.StatusFollow = true
	}
	utils.Response(ctx, w, true, st, stFoll)
}

func (u *ctrl) HandlerPostFollow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	user, msg, code, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, code, msg)
		return
	}

	ctx, followSt, msg, st, err := u.uc.GetFollowStatus(ctx, user.ID, id)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	if (followSt != nil) && (followSt.UserFollowID == 0) {
		ctx, msg, st, err = u.uc.PostFollow(ctx, user.ID, id)
		if err != nil {
			ctx = logger.Logf(ctx, "Do Follow error() => %v", err)
			utils.Response(ctx, w, false, st, msg)
			return
		}
	} else if followSt.Deleted == 1 {
		ctx, msg, st, err = u.uc.PostRefollow(ctx, user.ID, id)
		if err != nil {
			ctx = logger.Logf(ctx, "Do ReFollow error() => %v", err)
			utils.Response(ctx, w, false, st, msg)
			return
		}
	} else {
		ctx, msg, st, err = u.uc.PostUnfollow(ctx, user.ID, id)
		if err != nil {
			ctx = logger.Logf(ctx, "Do Unfollow error() => %v", err)
			utils.Response(ctx, w, false, st, msg)
			return
		}
	}

	utils.Response(ctx, w, true, st, msg)
}

func (u *ctrl) HandlerGetCountFollowing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	ctx, res, msg, st, err := u.uc.GetCountFollowing(ctx, id)
	if err != nil {
		ctx = logger.Logf(ctx, "Count Following error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetCountFollower(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	ctx, res, msg, st, err := u.uc.GetCountFollower(ctx, id)
	if err != nil {
		ctx = logger.Logf(ctx, "Count Follower error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetListFollower(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, res, msg, st, err := u.uc.GetListFollower(ctx, user.ID, id)
	if err != nil {
		ctx = logger.Logf(ctx, "List Follower error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetListFollowing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		ctx = logger.Logf(ctx, "Error on get request() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, res, msg, st, err := u.uc.GetListFollowing(ctx, user.ID, id)
	if err != nil {
		ctx = logger.Logf(ctx, "List Following error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}
