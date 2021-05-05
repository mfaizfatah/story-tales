package models

type Likes struct {
	// the id for idLikes
	//
	// readOnly: true
	// required: false
	// min: 1
	ID int `json:"id" gorm:"column:id"`
	// Like for story
	//
	// required: true
	// min: 1
	Like int `json:"like" gorm:"column:like"`
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

type ResponseCheckLikes struct {
	CheckLikes bool `json:"checkLikes"`
}
