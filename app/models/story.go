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
