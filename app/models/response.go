package models

// ResponseLogin is reponse for frontend cms
type ResponseLogin struct {
	Token    TokenResponse `json:"token,omitempty"`
	Roles    string        `json:"roles,omitempty"`
	IDAuthor int           `json:"id_author,omitempty"`
	Name     string        `json:"name,omitempty"`
	Email    string        `json:"email,omitempty"`
	Message  string        `json:"message,omitempty"`
}
