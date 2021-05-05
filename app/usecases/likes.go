package usecases

import (
	"context"
	"fmt"
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

func (r *uc) GetLikes(ctx context.Context, storyID, episodeID, userID int) (context.Context, *models.ResponseCheckLikes, string, int, error) {
	var (
		msg   string
		err   error
		likes = new(models.Likes)
		data  = new(models.ResponseCheckLikes)
	)

	err = r.query.FindOne(tableLikes, likes, "id_story = ? AND id_episodes = ? AND id_users = ?", "*", storyID, episodeID, userID)

	data.CheckLikes = false

	if err == nil {
		data.CheckLikes = true
	}

	if err != nil && fmt.Sprint(err) != "record not found" {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil
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
