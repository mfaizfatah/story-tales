// Package classification Story Tales API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - multipart/form-data
//
//     Produces:
//     - application/json
// swagger:meta
package controllers

import (
	"github.com/mfaizfatah/story-tales/app/models"
)

// Data structure representing a single story
// swagger:response postResponse
type postResponseWrapper struct {
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

// swagger:parameters createStory
type postStorysWrapper struct {
	/*
		The id of the product for which the operation relates
		in: body
		required: true
	*/
	Body models.Story
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
		in: path
		required: true
	*/

	StoryID int `json:"storyID"`
}
