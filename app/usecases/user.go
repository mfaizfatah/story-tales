package usecases

import (
	"context"
	"log"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableAuthorView = "authorProfileView"
	tableAuthor     = "author"
)

func (r *uc) GetExistAuthor(ctx context.Context, req *models.AuthorNickName) (context.Context, *models.AuthorNickName, string, int, error) {
	var (
		res    = new(models.AuthorNickName)
		author = new(models.AuthorNickName)
		msg    string
		err    error
	)

	err = r.query.FindOne(tableAuthorView, author, "author_nickname = ?", "id, author_nickname", req.AuthorNickName)
	if author.AuthorNickName != "" && err == nil {
		return ctx, nil, ErrAlreadyUserName, http.StatusConflict, repository.ErrConflict
	}
	res.AuthorNickName = req.AuthorNickName
	return ctx, res, msg, http.StatusCreated, nil
}

func (r *uc) GetAuthorProfile(ctx context.Context, authorID int) (context.Context, *models.AuthorProfile, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetAuthorProfile(authorID, tableAuthorView)
	log.Printf("msg: %v", data)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) UpdateAuthor(ctx context.Context, req *models.AuthorData, authorID int) (context.Context, string, int, error) {
	var (
		authorData = new(models.AuthorData)
		msg        string
		err        error
	)

	err = r.query.UpdateData(tableAuthor, authorData, "id_user = ?", req, authorID)
	if err != nil {
		return ctx, ErrNotFound, http.StatusInternalServerError, err
	}
	return ctx, msg, http.StatusCreated, err
}
