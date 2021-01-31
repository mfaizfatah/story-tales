package models

// Story model data
type Story struct {
	ID              int               `json:"id" gorm:"column:id"`
	Title           string            `json:"title" gorm:"column:title"`
	Season          string            `json:"season" gorm:"column:season"`
	Sinopsis        string            `json:"sinopsis" gorm:"column:sinopsis"`
	Images          string            `json:"images" gorm:"column:images"`
	FlagOnGoing     int               `json:"flagOnGoing" gorm:"column:flag_ongoing"`
	FlagCommment    int               `json:"flagComment" gorm:"column:flag_comment"`
	IDAuthor        int               `json:"idAuthor" gorm:"column:id_author"`
	Episode         []Episode         `json:"episode" gorm:"foreignKey:id_story;references:ID"`
	Episodes_Detail []Episodes_Detail `json:"episodeDetail"gorm:"foreignKey:id_story;references:ID"`
}

type Episode struct {
	ID         int    `json:"id" gorm:"column:id"`
	ID_Story   int    `json:"idStory" gorm:"column:id_story"`
	Eps_Number int    `json:"epsNumber" gorm:"column:eps_number"`
	Eps_Title  string `json:"epsTitle" gorm:"column:eps_title"`
}

type Episodes_Detail struct {
	ID         int    `json:"id" gorm:"column:id"`
	ID_Story   int    `json:"idStory" gorm:"column:id_story"`
	ID_Episode int    `json:"idEpisode" gorm:"column:id_episodes"`
	Page       int    `json:"page" gorm:"column:page"`
	Schedule   string `json:"schedule" gorm:"column:schedule"`
	Images     string `json:"images" gorm:"column:images"`
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

//GetDetail
type ResponseDetailEpisode struct {
	ID         int      `json:"id"`
	Eps_Number int      `json:"eps_number"`
	Eps_Title  string   `json:"eps_title"`
	Detail     []Detail `json:"detail"`
}

type Detail struct {
	ID       int    `json:"id"`
	Page     int    `json:"page"`
	Schedule string `json:"schedule"`
	Images   string `json:"images"`
}
