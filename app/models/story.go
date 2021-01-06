package models

// Story model data
type Story struct {
	ID           int    `json:"id" gorm:"column:id"`
	Title        string `json:"title" gorm:"column:title"`
	Season       string `json:"season" gorm:"column:season"`
	Sinopsis     string `json:"sinopsis" gorm:"column:sinopsis"`
	Images       string `json:"images" gorm:"column:images"`
	FlagOnGoing  int    `json:"flagOnGoing" gorm:"column:flag_ongoing"`
	FlagCommment int    `json:"flagComment" gorm:"column:flag_comment"`
	IDAuthor     int    `json:"idAuthor" gorm:"column:id_author"`
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
	IDAuthor     int           `json:"idAuthor"`
	ListEpisode  []ListEpisode `json:"listEpisode"`
}

//ListEpisode
type ListEpisode struct {
	ID         int    `json:"id"`
	Eps_Number int    `json:"eps_number"`
	Eps_Title  string `json:"eps_title"`
}
