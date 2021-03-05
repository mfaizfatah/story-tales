package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

func (r *uc) GetFavoriteStory(ctx context.Context, userid int) (context.Context, []models.ResponseFavoriteStory, string, int, error) {

	var (
		msg string
		err error
	)

	data, err := r.query.FindFavoriteStory(tableStory, userid)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) PostFavoriteStory(ctx context.Context, req *models.PostFavoriteStory, userid int) (context.Context, string, int, error) {
	var (
		fav = new(models.PostFavoriteStory)
		msg = "Favorite Berhasil Ditambahkan"
		err error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	fav = req
	fav.IDUsers = userid
	err = r.query.Insert("user_favorite", fav)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) DeleteFavoriteStory(ctx context.Context, storyid, userid int) (context.Context, string, int, error) {
	var (
		msg string
		err error
	)

	msg = "Favorite Berhasil Dihapus"

	err = r.query.DeleteFavorite(storyid, userid)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusOK, err
}
