package usecases

import (
	"context"
	"log"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableuserFollow = "user_follow"
)

func (r *uc) GetCountFollowing(ctx context.Context, id int) (context.Context, *models.UserCountFollowing, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetCountFollowing(id)
	log.Printf("msg: %v", data)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetCountFollower(ctx context.Context, id int) (context.Context, *models.UserCountFollower, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetCountFollower(id)
	log.Printf("msg: %v", data)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}
