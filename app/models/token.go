package models

import "time"

// TokenValue is model data for token value
type TokenValue struct {
	IDUser   int    `json:"id_user"`
	FCMToken string `json:"fcm_token"`
}

// TokenResponse is model data for response token
type TokenResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ExpiredIn string `json:"expired_in"`
}

// Session ...
type Session struct {
	IDUser  string         `json:"id_user" bson:"id_user"`
	Session []SessionToken `json:"session" bson:"session"`
	// Session  []string `json:"session" bson:"session"`
}

//SessionToken is func to get session
type SessionToken struct {
	Token     string    `json:"token" bson:"token"`
	FcmToken  string    `json:"fcm_token" bson:"fcm_token"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
