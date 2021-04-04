package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableComment = "story_comment"
	viewComment  = "commentView"
)

func (r *uc) GetListComment(ctx context.Context, storyID, episodeID int) (context.Context, []models.CommentView, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindAllComment(viewComment, storyID, episodeID)

	if err != nil {
		return ctx, nil, ErrBadRequest, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusAccepted, nil

}

func (r *uc) PostComment(ctx context.Context, req *models.Comment, userID int) (context.Context, string, int, error) {
	var (
		comment = new(models.Comment)
		msg     string
		err     error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	comment = req
	comment.IDUser = userID
	err = r.query.Insert(tableComment, comment)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) DeleteComment(ctx context.Context, commentID, userID int) (context.Context, string, int, error) {
	var (
		comment = new(models.Comment)
		msg     string
		err     error
	)
	msg = "Comment Berhasil Dihapus"

	comment.ID = commentID
	comment.IDUser = userID
	err = r.query.Delete(tableComment, comment)

	if err != nil {
		return ctx, ErrServer, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusOK, err
}
