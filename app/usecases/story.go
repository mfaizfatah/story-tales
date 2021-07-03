package usecases

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	tableStory     = "story"
	StoryGenreView = "storyGenreView"
	detailStory    = "detailStoryOneView"
)

func (r *uc) PostStory(ctx context.Context, req *models.Story, userid int) (context.Context, string, int, error) {
	var (
		story = new(models.Story)
		msg   string
		err   error
	)

	if req == nil {
		return ctx, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	story = req
	story.IDAuthor = userid
	err = r.query.Insert(tableStory, story)
	if err != nil {
		return ctx, ErrCreated, http.StatusInternalServerError, err
	}

	// for _, vals := range req.Genre {
	// 	StoryGenre := new(models.Story_Genre)
	// 	StoryGenre.IDStory = story.ID
	// 	StoryGenre.IDGenre = vals.IDGenre
	// 	log.Println(StoryGenre)
	// 	err = r.query.Insert("story_genres", StoryGenre)
	// 	if err != nil {
	// 		return ctx, ErrCreated, http.StatusInternalServerError, err
	// 	}
	// }

	msg = fmt.Sprintf("Successfull insert story with id %v", story.ID)
	return ctx, msg, http.StatusCreated, err
}

func (r *uc) GetOneStory(ctx context.Context, storyID int) (context.Context, *models.ResponseOneStory, string, int, error) {
	var (
		msg       string
		err       error
		totalLike int
	)

	data, err := r.query.FindGetOneStory(storyID)

	for i := 0; i < len(data.ListEpisode); i++ {
		totalLike += data.ListEpisode[i].Like
	}
	data.TotalLike = totalLike

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetDetailEpisode(ctx context.Context, storyID, episodeID int) (context.Context, *models.ResponseDetailEpisode, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindGetDetailEpisode(storyID, episodeID)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetAllStory(ctx context.Context) (context.Context, []models.ResponseAllStory, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindAllStory(detailStory)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetRekomendasiStory(ctx context.Context) (context.Context, []models.ResponseRekomenStory, string, int, error) {
	var (
		msg string
		err error
	)

	data, err := r.query.FindRekomendasiStory(tableStory)

	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, data, msg, http.StatusOK, nil

}

func (r *uc) GetStoryGenre(ctx context.Context) (context.Context, []models.ResponseStoryGenre, string, int, error) {
	var (
		res        []models.ResponseStoryGenre
		msg        string
		code       = http.StatusOK
		storyGenre []models.StoryGenreView
	)

	err := r.query.DBFindAll(StoryGenreView, &storyGenre, "deleted = ?", "id_story, genre, title, images, sinopsis, nickname", 0)
	if err != nil {
		return ctx, nil, "data_not_found", http.StatusNotFound, repository.ErrRecordNotFound
	}

	idStory := make(map[string][]int)
	title := make(map[string][]string)
	image := make(map[string][]string)
	sinopsis := make(map[string][]string)
	nickname := make(map[string][]string)
	for _, vals := range storyGenre {
		idStory[vals.Genre] = append(idStory[vals.Genre], vals.IDStory)
		title[vals.Genre] = append(title[vals.Genre], vals.Title)
		image[vals.Genre] = append(image[vals.Genre], vals.Images)
		sinopsis[vals.Genre] = append(sinopsis[vals.Genre], vals.Sinopsis)
		nickname[vals.Genre] = append(nickname[vals.Genre], vals.AuthorNickName)
	}

	for i, j := range idStory {
		result := models.ResponseStoryGenre{}
		result.Genre = i

		for k, l := range j {
			storyView := models.StoryGenreView{}
			storyView.IDStory = l
			storyView.Title = title[i][k]
			storyView.Images = image[i][k]
			storyView.Sinopsis = sinopsis[i][k]
			storyView.AuthorNickName = nickname[i][k]

			result.Story = append(result.Story, storyView)
		}
		res = append(res, result)
	}

	return ctx, res, msg, code, nil
}

func (r *uc) GetAuthorStory(ctx context.Context, authorID int) (context.Context, []models.ResponseAllStory, string, int, error) {
	var (
		msg string
		err error
	)
	data, err := r.query.FindAuthorStory(detailStory, authorID)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}
	return ctx, data, msg, http.StatusOK, nil
}

func (r *uc) GetAllGenre(ctx context.Context) (context.Context, []models.Genre, string, int, error) {
	var (
		msg string
		err error
		i   []models.Genre
	)

	err = r.query.DBFindAll("genre", &i, "", "id, genre")
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	return ctx, i, msg, http.StatusOK, nil

}
