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
	likeComment  = "story_comment_like"
)

func (r *uc) GetLikeCommentStatus(ctx context.Context, userID, commentID int) (context.Context, *models.CommentLike, string, int, error) {
	var (
		msg string
		err error
	)
	data, err := r.query.FindCommentLikeSt(commentID, userID)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil
}

func (r *uc) PostReLikeComment(ctx context.Context, userID, commentID int) (context.Context, string, int, error) {
	var (
		commentLike = new(models.CommentLike)
		msg         string
		err         error
	)
	data := make(map[string]interface{})
	data["deleted"] = 0

	err = r.query.UpdateWhere(likeComment, commentLike, "id_users = ? AND id_comment = ?", data, userID, commentID)
	if err != nil {
		return ctx, ErrNotFound, http.StatusInternalServerError, err
	}
	return ctx, msg, http.StatusCreated, err
}

func (r *uc) PostUnLikeComment(ctx context.Context, userID, commentID int) (context.Context, string, int, error) {
	var (
		commentLike = new(models.CommentLike)
		msg         string
		err         error
	)
	data := make(map[string]interface{})
	data["deleted"] = 1

	err = r.query.UpdateWhere(likeComment, commentLike, "id_users = ? AND id_comment = ?", data, userID, commentID)
	if err != nil {
		return ctx, ErrNotFound, http.StatusInternalServerError, err
	}
	return ctx, msg, http.StatusCreated, err
}

func (r *uc) PostLikeComment(ctx context.Context, userID, commentID int) (context.Context, string, int, error) {
	var (
		commentLike = new(models.CommentLike)
		msg         string
		err         error
	)
	commentLike.IDUser = userID
	commentLike.IDComment = commentID
	err = r.query.Insert(likeComment, commentLike)
	if err != nil {
		return ctx, ErrCreated, http.StatusInternalServerError, err
	}
	return ctx, msg, http.StatusCreated, err
}

func (r *uc) GetListTopComment(ctx context.Context, storyID, episodeID int) (context.Context, []models.CommentView, string, int, error) {
	var (
		msg string
		err error
	)
	data, err := r.query.FindTopComment(viewComment, storyID, episodeID)
	if err != nil {
		return ctx, nil, ErrBadRequest, http.StatusNotFound, repository.ErrRecordNotFound
	}
	return ctx, data, msg, http.StatusAccepted, nil
}

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

func (r *uc) GetMyComment(ctx context.Context, userID int) (context.Context, []models.CommentView, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindMyComment(viewComment, userID)

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
