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

func (r *uc) PostStory(ctx context.Context, req *models.Story, userid int) (context.Context, string, int, error) {
	var (
		story = new(models.Story)
		msg   string
		err   error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	story = req
	story.IDAuthor = userid
	err = r.query.Insert(tableStory, story)

	if err != nil {
		return ctx, ErrCreated, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) GetOneStory(ctx context.Context, storyID int) (context.Context, *models.ResponseOneStory, string, int, error) {
	var (
		msg       string
		err       error
		totalLike int
	)

	data, err := r.query.FindGetOneStory(storyID)

	for i := 0; i < len(data.ListEpisode); i++ {
		totalLike += data.ListEpisode[i].Like
	}
	data.TotalLike = totalLike

	log.Printf("msg: %v", data)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetDetailEpisode(ctx context.Context, storyID, episodeID int) (context.Context, *models.ResponseDetailEpisode, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetDetailEpisode(storyID, episodeID)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetAllStory(ctx context.Context) (context.Context, []models.ResponseAllStory, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindAllStory(tableStory)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetRekomendasiStory(ctx context.Context) (context.Context, []models.ResponseRekomenStory, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindRekomendasiStory(tableStory)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}
