package models

import "time"

// StatusFollow is model for user statusfollow
type StatusFollow struct {
	StatusFollow bool `json:"statusfollow"`
}

// UserFollow is model for get user follow status
type UserFollow struct {
	UserFollowID    int `json:"userFollowId" gorm:"column:userfollow_id"`
	UserFollowingID int `json:"userFollowingId" gorm:"column:userfollowing_id"`
	Deleted         int `json:"deleted" gorm:"column:deleted"`
}

// UserCountFollowing is model for count user following
type UserCountFollowing struct {
	Count int64 `json:"following"`
}

// UserCountFollower is model for count user following
type UserCountFollower struct {
	Count int64 `json:"follower"`
}

// ListFollower is model for view user follower
type ListFollower struct {
	UserID      string    `json:"userId" gorm:"column:userfollow_id"`
	UserName    string    `json:"userName" gorm:"column:username"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	DateUpdated time.Time `json:"dateUpdated" gorm:"column:date_updated"`
}

// ListFollowing is model for view user following
type ListFollowing struct {
	UserID      string    `json:"userId" gorm:"column:userfollowing_id"`
	UserName    string    `json:"userName" gorm:"column:username"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	DateUpdated time.Time `json:"dateUpdated" gorm:"column:date_updated"`
}
