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

func (r *uc) PostRefollow(ctx context.Context, userID, id int) (context.Context, string, int, error) {
	var (
		usrFollow = new(models.UserFollow)
		msg       string
		err       error
	)
	data := make(map[string]interface{})
	data["deleted"] = 0

	err = r.query.UpdateWhere(tableuserFollow, usrFollow, "userfollow_id = ? AND userfollowing_id = ?", data, userID, id)
	if err != nil {
		return ctx, ErrNotFound, http.StatusInternalServerError, err
	}
	return ctx, msg, http.StatusCreated, err
}

func (r *uc) PostUnfollow(ctx context.Context, userID, id int) (context.Context, string, int, error) {
	var (
		usrFollow = new(models.UserFollow)
		msg       string
		err       error
	)
	data := make(map[string]interface{})
	data["deleted"] = 1

	err = r.query.UpdateWhere(tableuserFollow, usrFollow, "userfollow_id = ? AND userfollowing_id = ?", data, userID, id)
	if err != nil {
		return ctx, ErrNotFound, http.StatusInternalServerError, err
	}
	return ctx, msg, http.StatusCreated, err
}

func (r *uc) PostFollow(ctx context.Context, userID, id int) (context.Context, string, int, error) {
	var (
		usrFollow = new(models.UserFollow)
		msg       string
		err       error
	)

	usrFollow.UserFollowID = userID
	usrFollow.UserFollowingID = id
	err = r.query.Insert(tableuserFollow, usrFollow)

	if err != nil {
		return ctx, ErrCreated, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) GetFollowStatus(ctx context.Context, userID, id int) (context.Context, *models.UserFollow, string, int, error) {
	var (
		msg string
		err error
	)
	data, err := r.query.FindFollowSt(userID, id)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil
}

func (r *uc) GetCountFollowing(ctx context.Context, id int) (context.Context, *models.UserCountFollowing, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetCountFollowing(id)
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

func (r *uc) GetListFollower(ctx context.Context, userID, id int) (context.Context, []models.ListFollower, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindListFollower(id)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusAccepted, nil

}

func (r *uc) GetListFollowing(ctx context.Context, userID, id int) (context.Context, []models.ListFollowing, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindListFollowing(id)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusAccepted, nil

}
