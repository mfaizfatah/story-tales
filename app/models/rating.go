package models

type Rating struct {
	ID      int `json:"id" gorm:"column:id"`
	Rating  int `json:"rating" gorm:"column:rating"`
	IDStory int `json:"id_story" gorm:"id_story"`
	IDUser  int `json:"id_users" gorm:"column:id_users"`
}
