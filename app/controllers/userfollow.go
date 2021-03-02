package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerPostFollow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	ctx, msg, st, err := u.uc.PostFollow(ctx, id)
	if err != nil {
		ctx = logger.Logf(ctx, "Do Follow error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
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

	ctx, res, msg, st, err := u.uc.GetListFollower(ctx, id)
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

	ctx, res, msg, st, err := u.uc.GetListFollower(ctx, id)
	if err != nil {
		ctx = logger.Logf(ctx, "List Following error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}
