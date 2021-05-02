package usecases

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

func (r *uc) UploadImages(ctx context.Context, story *models.Story, userid int, file multipart.File, fileHeader *multipart.FileHeader) (context.Context, int, error) {
	var (
		err           error
		dir           = os.Getenv("IMAGES_DIRECTORY")
		baseUriImages = os.Getenv("IMAGES_URI")
	)

	if story == nil {
		return ctx, http.StatusBadRequest, repository.ErrBadRequest
	}

	// image segment
	title := strings.ToLower(strings.ReplaceAll(story.Title, " ", "_"))

	// /data/sftpdigi/upload/<idUser>/<title>/filename
	dirpath := filepath.Join(dir, strconv.Itoa(userid), title)
	if _, err = os.Stat(dirpath); os.IsNotExist(err) {
		err = os.MkdirAll(dirpath, 0755)
	}

	fileLocation := filepath.Join(dirpath, fileHeader.Filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return ctx, http.StatusInternalServerError, err
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, file); err != nil {
		return ctx, http.StatusInternalServerError, err
	}
	ctx = logger.Logf(ctx, "file() => %v", fileLocation)
	story.Images = fmt.Sprintf("%v/%v/%v/%v", baseUriImages, strconv.Itoa(userid), title, fileHeader.Filename)

	return ctx, http.StatusOK, err
}
