package usecases

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

func (r *uc) UploadImages(ctx context.Context, story *models.Story, userid int, file multipart.File, fileHeader *multipart.FileHeader) (context.Context, int, error) {
	var (
		err error
		// dir           = os.Getenv("IMAGES_DIRECTORY")
		baseUriImages = os.Getenv("IMAGES_URI")
	)

	if story == nil {
		return ctx, http.StatusBadRequest, repository.ErrBadRequest
	}

	// image segment
	title := strings.ToLower(strings.ReplaceAll(story.Title, " ", "_"))

	// ctx, fileLoc, err := UploadToFtpProccess(userid, story, file, fileHeader)

	story.Images = fmt.Sprintf("%v/%v/%v/%v", baseUriImages, strconv.Itoa(userid), title, fileHeader.Filename)

	return ctx, http.StatusOK, err
}
