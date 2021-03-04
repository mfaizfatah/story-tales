package models

type Likes struct {
	ID         int `json:"id" gorm:"column:id"`
	Like       int `json:"like" gorm:"column:like"`
	IDStory    int `json:"id_story" gorm:"id_story"`
	IDEpisodes int `json:"id_episodes" gorm:"column:id_episodes"`
	IDUser     int `json:"id_users" gorm:"column:id_users"`
}
