package models

import "time"

// Story defines the structure for an API Story
// swagger:model Story
type Story struct {
	// the id for the product
	//
	// readOnly: true
	// required: false
	// min: 1
	ID int `json:"id" gorm:"column:id"` // Unique identifier for story

	// the name for this story
	//
	// required: true
	// max length: 255
	Title string `json:"title" gorm:"column:title"`
	// season for story
	//
	// required: true
	// max length: 255
	Season string `json:"season" gorm:"column:season"`
	// Sinopsis Story
	//
	// required: true
	// max length:10000
	Sinopsis string `json:"sinopsis" gorm:"column:sinopsis"`
	// images Story
	//
	// required: true
	// max length:10000
	Images string `json:"images" gorm:"column:images"`
	// images Story
	//
	// required: true
	// max length:1
	FlagOnGoing int `json:"flagOnGoing" gorm:"column:flag_ongoing"`
	// images Story
	//
	// required: true
	// max length:1
	FlagCommment int `json:"flagComment" gorm:"column:flag_comment"`
	// the id for the author
	//
	// readOnly: true
	// required: false
	// min: 1
	IDAuthor int           `json:"idAuthor" gorm:"column:id_author"` // Unique identifier for author
	Episode  []Episode     `json:"episode" gorm:"foreignKey:id_story;references:ID"`
	Genre    []Story_Genre `json:"genre" gorm:"foreignKey:id_story;references:ID"`
}

type Episode struct {
	// the id for episode
	//
	// readOnly: true
	// required: false
	// min: 1
	ID int `json:"id" gorm:"column:id"`
	// the id for story
	//
	// readOnly: true
	// required: false
	// min: 1
	ID_Story int `json:"idStory" gorm:"column:id_story"`
	// number episode
	//
	// required: true
	// min: 1
	Eps_Number int `json:"epsNumber" gorm:"column:eps_number"`
	// title of episode
	//
	// required: true
	// max length: 255
	Eps_Title string `json:"epsTitle" gorm:"column:eps_title"`
	// images episode
	//
	// required: true
	// max length: 255
	Images_Eps      string            `json:"imagesEps" gorm:"column:images_eps"`
	Episodes_Detail []Episodes_Detail `json:"episodeDetail" gorm:"foreignKey:id_episodes;references:ID"`
}

type Episodes_Detail struct {
	// the id for episodeDetail
	//
	// readOnly: true
	// required: false
	// min: 1
	ID int `json:"id" gorm:"column:id"`
	// the id for episode
	//
	// required: true
	// min: 1
	ID_Episode int `json:"idEpisode" gorm:"column:id_episodes"`
	// page number episode
	//
	// required: true
	// min: 1
	Page int `json:"page" gorm:"column:page"`
	// schedule release date
	//
	// required: true
	// min: 1
	Schedule string `json:"schedule" gorm:"column:schedule"`
	// images episode
	//
	// required: true
	// min: 1
	Images string `json:"images" gorm:"column:images"`
}

// Response RekomendasiStory..
type ResponseRekomenStory struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Images string `json:"images"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

// Response All..
type ResponseAllStory struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Season       string `json:"season"`
	Images       string `json:"images"`
	FlagOnGoing  int    `json:"flagOnGoing"`
	FlagCommment int    `json:"flagComment"`
	IDAuthor     int    `json:"idAuthor"`
}

// Response Get One..
type ResponseOneStory struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	Sinopsis     string        `json:"sinopsis"`
	Season       string        `json:"season"`
	Images       string        `json:"images"`
	FlagOnGoing  int           `json:"flagOnGoing"`
	FlagCommment int           `json:"flagComment"`
	ID_Author    int           `json:"id_author"`
	Author       string        `json:"author"`
	Genre        string        `json:"genre"`
	TotalLike    int           `json:"totalLike"`
	Rating       float64       `json:"rating"`
	Publish_Date time.Time     `json:"publish_story_date"`
	ListEpisode  []ListEpisode `json:"listEpisode"`
}

//ListEpisode
type ListEpisode struct {
	ID             int       `json:"id"`
	Like           int       `json:"like"`
	Images_Eps     string    `json:"images_eps"`
	Eps_Number     int       `json:"eps_number"`
	Eps_Title      string    `json:"eps_title"`
	Publish_Status int       `json:"publish_episode_status"`
	Publish_Date   time.Time `json:"publish_episode_date"`
}

//GetDetail
type ResponseDetailEpisode struct {
	ID         int      `json:"id"`
	Eps_Number int      `json:"eps_number"`
	Eps_Title  string   `json:"eps_title"`
	ID_Story   int      `json:"id_story"`
	Likes      int      `json:"like"`
	Detail     []Detail `json:"detail"`
}

type Detail struct {
	ID       int    `json:"id"`
	Page     int    `json:"page"`
	Schedule string `json:"schedule"`
	Images   string `json:"images"`
}

type StoryGenreView struct {
	IDStory        int    `json:"id_story"`
	Genre          string `json:"genre,omitempty"`
	Title          string `json:"title"`
	Images         string `json:"images"`
	Sinopsis       string `json:"sinopsis"`
	AuthorNickName string `json:"nickname" gorm:"column:nickname"`
}

type ResponseStoryGenre struct {
	Genre string           `json:"Genre"`
	Story []StoryGenreView `json:"story"`
}

type Genre struct {
	ID    int    `json:"id" gorm:"column:id"`
	Genre string `json:"Genre" gorm:"column:genre"`
}

type Story_Genre struct {
	ID      int `json:"id" gorm:"column:id"`
	IDStory int `json:"idStory" gorm:"column:id_story"`
	IDGenre int `json:"idGenre" gorm:"column:id_genre"`
}
