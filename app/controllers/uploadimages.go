package controllers

import (
	"net/http"
	"strings"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	path := r.FormValue("path")

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	// image segment
	if err := r.ParseMultipartForm(1024); err != nil {
		ctx = logger.Logf(ctx, "upload error() => %v", err)
		utils.Response(ctx, w, false, http.StatusInternalServerError, err.Error())
		return
	}
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		ctx = logger.Logf(ctx, "upload error() => %v", err)
		utils.Response(ctx, w, false, http.StatusInternalServerError, err.Error())
		return
	}
	defer uploadedFile.Close()

	path = strings.ToLower(strings.ReplaceAll(path, " ", "_"))
	ctx, res, err := u.uc.UploadToFtpProccess(ctx, user.ID, path, uploadedFile, handler)
	if err != nil {
		ctx = logger.Logf(ctx, "upload error() => %v", err)
		utils.Response(ctx, w, false, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Response(ctx, w, true, st, res)
}
