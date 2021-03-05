package models

type Rating struct {
	// the id for idRating
	//
	// readOnly: true
	// required: false
	// min: 1
	ID int `json:"id" gorm:"column:id"`
	// rating for story
	//
	// required: true
	// min: 1
	Rating int `json:"rating" gorm:"column:rating"`
	// the id for story
	//
	// required: true
	// min: 1
	IDStory int `json:"id_story" gorm:"id_story"`
	// the id for idUser
	//
	// readOnly: true
	// required: false
	// min: 1
	IDUser int `json:"id_users" gorm:"column:id_users"`
}
