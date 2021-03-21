package models

import "time"

type SearchModel struct {
	Title     string    `json:"title"`
	Sinopsis  string    `json:"sinopsis"`
	CreatedAt time.Time `json:"created_at"`
}
