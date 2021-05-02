package controllers

import (
	"net/http"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var s models.Story

	s.Title = r.FormValue("title")

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	// image segment
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	ctx, st, err = u.uc.UploadImages(ctx, &s, user.ID, uploadedFile, handler)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, err.Error())
		return
	}

	utils.Response(ctx, w, true, st, s.Images)
}
