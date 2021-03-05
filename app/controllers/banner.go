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

func (u *ctrl) HandlerCreateBanner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s models.BannerReq

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, msg, st, err := u.uc.CreateBanner(ctx, &s)
	if err != nil {
		ctx = logger.Logf(ctx, "Banner error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

func (u *ctrl) HandlerGetBannerDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	ctx, res, msg, st, err := u.uc.GetBannerDtl(ctx, id)
	if err != nil {
		ctx = logger.Logf(ctx, "Banner Detail error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetListBannerThumb(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, res, msg, st, err := u.uc.GetListBannerThumb(ctx)
	if err != nil {
		ctx = logger.Logf(ctx, "List Banner error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}
