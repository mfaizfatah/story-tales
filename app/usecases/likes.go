package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableLikes = "likes"
)

func (r *uc) PostLikes(ctx context.Context, req *models.Likes, userid int) (context.Context, string, int, error) {
	var (
		likes = new(models.Likes)
		msg   string
		err   error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	likes = req
	likes.IDUser = userid
	err = r.query.Insert(tableLikes, likes)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) DeleteLikes(ctx context.Context, storyid, episodeid, userid int) (context.Context, string, int, error) {
	var (
		msg string
		err error
	)

	msg = "Likes Berhasil Dihapus"

	err = r.query.DeleteLikes(storyid, episodeid, userid)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusOK, err
}

func (r *uc) GetMyLikes(ctx context.Context, userID int) (context.Context, []models.Likes, string, int, error) {
	var (
		msg string
		err error
	)
	data, err := r.query.FindMyLike(tableLikes, userID)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}
	return ctx, data, msg, http.StatusOK, nil
}
