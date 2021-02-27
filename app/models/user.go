package models

import "time"

// User is model for user data
type User struct {
	ID                int         `json:"id_user" gorm:"column:id"`
	Username          string      `json:"username" gorm:"column:username"`
	Email             string      `json:"email" gorm:"column:email"`
	Password          string      `json:"password" gorm:"column:password"`
	Name              string      `json:"name" gorm:"column:name"`
	Telp              string      `json:"telp" gorm:"column:telp"`
	DateOfBirth       interface{} `json:"dateOfBirth" gorm:"column:date_of_birth"`
	IDRole            int         `json:"idRole" gorm:"column:id_role"`
	EmailVerification int         `json:"emailVerify" gorm:"column:email_verify"`
}

// ForgotPass is mode for request forgot pass
type ForgotPass struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

// UserFollow is model for get user follow status
type UserFollow struct {
	UserFollowID    string    `json:"userFollowId" gorm:"column:userfollow_id"`
	UserFollowingID string    `json:"userFollowingId" gorm:"column:userfollowing_id"`
	DateUpdated     time.Time `json:"dateUpdated" gorm:"column:date_updated"`
}

// UserCountFollowing is model for count user following
type UserCountFollowing struct {
	Count int64 `json:"following"`
}

// UserCountFollower is model for count user following
type UserCountFollower struct {
	Count int64 `json:"follower"`
}
