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

func (u *ctrl) HandlerGetExistAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var author models.AuthorNickName
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, res, msg, st, err := u.uc.GetExistAuthor(ctx, &author)
	if err != nil {
		ctx = logger.Logf(ctx, "Check ExistAuthor error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetAuthorProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authorID, err := strconv.Atoi(chi.URLParam(r, "authorID"))

	ctx, res, msg, st, err := u.uc.GetAuthorProfile(ctx, authorID)
	if err != nil {
		ctx = logger.Logf(ctx, "Get Author Profile error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerUpdateAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var authorData models.AuthorData

	err := json.NewDecoder(r.Body).Decode(&authorData)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.UpdateAuthor(ctx, &authorData, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Update Author Error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}
