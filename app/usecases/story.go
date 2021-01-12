package usecases

import (
	"context"
	"log"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableStory = "story"
)

func (r *uc) PostStory(ctx context.Context, req *models.Story) (context.Context, string, int, error) {
	var (
		story = new(models.Story)
		msg   string
		err   error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	story = req
	log.Printf("msg: %v", story)
	err = r.query.Insert(tableStory, story)

	if err != nil {
		return ctx, ErrCreated, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) GetOneStory(ctx context.Context, storyID int) (context.Context, *models.ResponseOneStory, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetOneStory(storyID)
	log.Printf("msg: %v", data)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusAccepted, nil

}

func (r *uc) GetAllStory(ctx context.Context) (context.Context, []models.ResponseAllStory, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindAll(tableStory)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusAccepted, nil

}
