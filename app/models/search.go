package models

import "time"

type SearchModel struct {
	IDStory   int       `json:"id_story"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
