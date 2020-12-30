package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableStory = "story"
)

func (r *uc) InsertStory(ctx context.Context, req *models.Story) (context.Context, string, int, error) {
	var (
		story = new(models.Story)
		msg   string
		err   error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	story = req
	err = r.query.Insert(tableStory, story)

	if err != nil {
		return ctx, ErrCreated, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) GetStory(ctx context.Context, storyID int) (context.Context, *models.ResponseStory, string, int, error) {
	var (
		res = new(models.ResponseStory)
		msg string
		err error
	)

	err = r.query.FindOne(tableStory, res, "id = ?", "id", storyID)
	if err == nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, res, msg, http.StatusAccepted, nil

}
