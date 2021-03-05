package controllers

import (
	"github.com/mfaizfatah/story-tales/app/models"
)

// Data structure representing error
// swagger:response postResponse
type postResponseWrapper struct {
	// in: body
	response struct {
		Status       string      `json:"status"`
		ErrorMessage string      `json:"error_message"`
		Data         interface{} `json:"data"`
	}
}

// Data structure representing error
// swagger:response errorResponse
type errorResponseWrapper struct {
	// in: body
	response struct {
		Status       string      `json:"status"`
		ErrorMessage string      `json:"error_message"`
		Data         interface{} `json:"data"`
	}
}

// Data structure representing a single story
// swagger:response getOneStoryResponse
type getOneStoryResponseWrapper struct {
	//Get One Story
	// in: body
	Body models.ResponseOneStory
}

// Data structure representing Rekomendasi story
// swagger:response getRekomenStoryResponse
type getRekomenStoryResponseWrapper struct {
	//Get Rekomen
	// in: body
	Body models.ResponseRekomenStory
}

// Data structure representing all story
// swagger:response getAllStoryResponse
type getAllStoryResponseWrapper struct {
	//Get All Story
	// in: body
	Body models.ResponseAllStory
}

// Data structure representing all story
// swagger:response getDetailEpisodeResponse
type getdDetailEpisodeWrapper struct {
	//Get Detail Episode
	// in: body
	Body models.ResponseDetailEpisode
}

// Data structure representing favorite story
// swagger:response getFavoriteResponse
type getFavoriteResponseWrapper struct {
	//Get Favorite Story
	// in: body
	Body models.ResponseFavoriteStory
}

// swagger:parameters postFavorite
type postFavoriteWrapper struct {
	/*
		The id of the product for which the operation relates
		in: body
		required: true
	*/
	Body models.PostFavoriteStory
}

// swagger:parameters deleteFavorite
type favoriteIDparamsWrapper struct {
	/*
		The id of the product for which the operation relates
		in: path
		required: true
	*/

	StoryID int `json:"storyID"`
}

// swagger:parameters postRating
type postRatingWrapper struct {
	/*
		The id of the product for which the operation relates
		in: body
		required: true
	*/
	Body models.Rating
}

// swagger:parameters deleteRating
type ratingIDparamsWrapper struct {
	/*
		The id of the product for which the operation relates
		in: path
		required: true
	*/

	StoryID int `json:"storyID"`
}

// swagger:parameters postLikes
type postLikesWrapper struct {
	/*
		The id of the product for which the operation relates
		in: body
		required: true
	*/
	Body models.Likes
}

// swagger:parameters createStory
type postStorysWrapper struct {
	/*
		The id of the product for which the operation relates
		in: body
		required: true
	*/
	Body models.Story
}

// swagger:parameters favoriteStory
type getFavoriteWrapper struct {
	/*
		The id of the product for which the operation relates
		in: body
		required: true
	*/
	Body models.ResponseFavoriteStory
}

// swagger:parameters detailEpisode
type episodeIDparamsWrapper struct {
	/*
		The id of the product for which the operation relates
		in: path
		required: true
	*/

	StoryID   int `json:"storyID"`
	EpisodeID int `json:"episodeID"`
}

// swagger:parameters oneStory
type storyIDParamsWrapper struct {
	/*
		The id of the product for which the operation relates

		In: path

		required: true
	*/

	StoryID int `json:"storyID"`
}
