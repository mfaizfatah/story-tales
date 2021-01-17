package usecases

import (
	"context"
	"log"
	"net/http"

	"github.com/go-shadow/moment"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableBanner = "banner"
)

func (r *uc) CreateBanner(ctx context.Context, req *models.Banner) (context.Context, string, int, error) {
	var (
		banner = new(models.Banner)
		msg    string
		err    error
	)

	if req.Title == "" {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	m := moment.New()
	currentDate := m.Now()
	createAt := currentDate.GetTime()
	validUntil := currentDate.Add("days", 1).GetTime()

	banner = req
	req.CreateAt = createAt
	req.ValidUntil = validUntil
	log.Printf("msg: %v", banner)
	err = r.query.Insert(tableBanner, banner)

	if err != nil {
		return ctx, ErrCreated, http.StatusInternalServerError, err
	}

	return ctx, msg, http.StatusCreated, err
}

func (r *uc) GetBannerDtl(ctx context.Context, id int) (context.Context, *models.BannerDetailRs, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetBanner(id)
	log.Printf("msg: %v", data)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetListBannerThumb(ctx context.Context) (context.Context, []models.ListBannerThumbRs, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindAllBanner(tableBanner)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusAccepted, nil

}
