package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableRating = "rating"
)

func (r *uc) PostRating(ctx context.Context, req *models.Rating, userid int) (context.Context, string, int, error) {
	var (
		rating = new(models.Rating)
		msg    string
		err    error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	rating = req
	rating.IDUser = userid
	err = r.query.Insert(tableRating, rating)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) DeleteRating(ctx context.Context, storyid, userid int) (context.Context, string, int, error) {
	var (
		msg string
		err error
	)

	err = r.query.DeleteRating(storyid, userid)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusOK, err
}
