package models

import "time"

type Comment struct {
	// the id for idComment
	//
	// readOnly: true
	// required: false
	// min: 1
	ID int `json:"id" gorm:"column:id"`
	// Comment for story
	//
	// required: true
	// min: 1
	Comment string `json:"comment" gorm:"column:comment"`
	// the id for story
	//
	// required: true
	// min: 1
	IDStory int `json:"id_story" gorm:"id_story"`
	// the id for episode
	//
	// required: true
	// min: 1
	IDEpisodes int `json:"id_episodes" gorm:"column:id_episodes"`
	// the id for idUsers
	//
	// readOnly: true
	// required: false
	// min: 1
	IDUser int `json:"id_users" gorm:"column:id_users"`
}

type CommentView struct {
	// the id for idComment
	//
	// readOnly: true
	// required: false
	// min: 1
	ID int `json:"id" gorm:"column:id"`
	// Comment for story
	//
	// required: true
	// min: 1
	Comment string `json:"comment" gorm:"column:comment"`
	// the id for story
	//
	// required: true
	// min: 1
	IDStory int `json:"id_story" gorm:"id_story"`
	// the id for episode
	//
	// required: true
	// min: 1
	IDEpisodes int `json:"id_episodes" gorm:"column:id_episodes"`
	// the id for idUsers
	//
	// readOnly: true
	// required: false
	// min: 1
	IDUser    int       `json:"id_users" gorm:"column:id_users"`
	UserName  string    `json:"username" gorm:"column:username"`
	Name      string    `json:"name" gorm:"column:name"`
	Avatar    string    `json:"avatar" gorm:"column:avatar"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
