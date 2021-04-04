package usecases

import (
	"log"

	"context"
	"net/http"

	"github.com/go-shadow/moment"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableBanner = "banner"
)

func (r *uc) CreateBanner(ctx context.Context, req *models.BannerReq) (context.Context, string, int, error) {
	var (
		banner = new(models.Banner)
		msg    string
		err    error
	)

	if req.Title == "" {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	currentDate := moment.New().Now()
	validUntil := currentDate.AddDays(req.DaysValid).GetTime()
	banner.Category = req.Category
	banner.Content = req.Content
	banner.Title = req.Title
	banner.DetailStatus = req.DetailStatus
	banner.DeepLink = req.DeepLink
	banner.URL = req.URL
	banner.Sequence = req.Sequence
	banner.ServiceID = req.ServiceID
	banner.Status = req.Status
	banner.ValidUntil = validUntil

	var imgLoc = "http://digisoul.id/images/"
	banner.Thumb = imgLoc + req.ThumbFile.Filename
	banner.Image = imgLoc + req.ImgFile.Filename

	log.Printf("WOOOOOOOYY: %v", banner.Thumb)
	log.Printf("ANNJAJAAY: %v", banner.Image)

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
