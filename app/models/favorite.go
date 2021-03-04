package models

// Response ResponseFavoriteStory..
type ResponseFavoriteStory struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Images string `json:"images"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

// Response ResponseFavoriteStory..
type PostFavoriteStory struct {
	ID      int `json:"id" gorm:"column:id"`
	IDStory int `json:"favorite_story" gorm:"column:favorite_story"`
	IDUsers int `json:"id_users" gorm:"column:id_users"`
}

// Response DeleteFavoriteStory..
type User_Favorite struct {
	ID      int `json:"id" gorm:"column:id"`
	IDStory int `json:"favorite_story" gorm:"column:favorite_story"`
	IDUsers int `json:"id_users" gorm:"column:id_users"`
}

func (User_Favorite) TableName() string {
	return "user_favorite"
}
