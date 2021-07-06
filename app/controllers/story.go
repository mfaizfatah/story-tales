package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

// swagger:route POST /story story createStory
// Return a list of story from the database
// responses:
//	200: postResponse
//	404: errorResponse
//
// ListOne handles GET requests and returns one story
func (u *ctrl) HandlerPostStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s models.Story

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, author, msg, st, err := u.uc.GetAuthorProfile(ctx, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.PostStory(ctx, &s, author.IDAuthor)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}

// swagger:route GET /story/{storyID} story oneStory
// Return a list of story from the database
// responses:
//	200: getOneStoryResponse
//	404: errorResponse
//
// ListOne handles GET requests and returns one story
func (u *ctrl) HandlerGetOneStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID, err := strconv.Atoi(chi.URLParam(r, "storyID"))

	ctx, res, msg, st, err := u.uc.GetOneStory(ctx, storyID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

// swagger:route GET /story/{storyID}/{episodeID} story detailEpisode
// Return a list of story from the database
//
// responses:
//	200: getDetailEpisodeResponse
//	404: errorResponse
//
// ListAll handles GET requests and returns all Story
func (u *ctrl) HandlerGetAllStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, res, msg, st, err := u.uc.GetAllStory(ctx)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

// swagger:route GET /story/rekomendasi story detailRekomen
// Return a list of story from the database
//
// responses:
//	200: getRekomenStoryResponse
//	404: errorResponse
//
// ListAll handles GET requests and returns recommend story
func (u *ctrl) HandlerGetRekomenStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, res, msg, st, err := u.uc.GetRekomendasiStory(ctx)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

// swagger:route GET /story story allStory
// Return a list of story from the database
//
// responses:
//	200: getAllStoryResponse
//	404: errorResponse
//
// ListAll handles GET requests and returns all Story
func (u *ctrl) HandlerGetDetailEpisode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	storyID, err := strconv.Atoi(chi.URLParam(r, "storyID"))
	episodeID, err := strconv.Atoi(chi.URLParam(r, "episodeID"))
	ctx, res, msg, st, err := u.uc.GetDetailEpisode(ctx, storyID, episodeID)
	if err != nil {
		ctx = logger.Logf(ctx, "Detail error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetAuthorStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authorID, err := strconv.Atoi(chi.URLParam(r, "authorID"))
	ctx, res, msg, st, err := u.uc.GetAuthorStory(ctx, authorID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetStoryAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, res, msg, st, err := u.uc.GetAuthorStory(ctx, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}
	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetStoryGenre(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, res, msg, st, err := u.uc.GetStoryGenre(ctx)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerGetAllGenre(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, res, msg, st, err := u.uc.GetAllGenre(ctx)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}

func (u *ctrl) HandlerUpdateStory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s models.Story

	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	user, msg, st, err := u.uc.GetUserFromToken(r)
	if err != nil {
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, author, msg, st, err := u.uc.GetAuthorProfile(ctx, user.ID)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	ctx, msg, st, err = u.uc.UpdateStory(ctx, &s, author.IDAuthor)
	if err != nil {
		ctx = logger.Logf(ctx, "Story error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, msg)
}
