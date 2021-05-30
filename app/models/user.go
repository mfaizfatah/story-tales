package models

// User is model for user data
type User struct {
	ID                int         `json:"id_user" gorm:"column:id"`
	IDAuthor          int         `json:"id_author" gorm:"-"`
	Username          string      `json:"username" gorm:"column:username"`
	Email             string      `json:"email" gorm:"column:email"`
	Password          string      `json:"password" gorm:"column:password"`
	Name              string      `json:"name" gorm:"column:name"`
	Telp              string      `json:"telp" gorm:"column:telp"`
	DateOfBirth       interface{} `json:"dateOfBirth" gorm:"column:date_of_birth"`
	Google            int         `json:"google" gorm:"column:google"`
	Avatar            string      `json:"avatar" gorm:"column:avatar"`
	IDRole            int         `json:"idRole" gorm:"column:id_role"`
	EmailVerification int         `json:"emailVerify" gorm:"column:email_verify"`
	User              string      `json:"user" gorm:"-"`
	FCMToken          string      `json:"fcm_token" gorm:"-"`
}

// Login is model for user data
type Login struct {
	ID                int    `json:"id_user" gorm:"column:id"`
	Username          string `json:"username" gorm:"column:username"`
	Email             string `json:"email" gorm:"column:email"`
	Password          string `json:"password" gorm:"column:password"`
	Google            int    `json:"google" gorm:"column:google"`
	EmailVerification int    `json:"emailVerify" gorm:"column:email_verify"`
	Role              string `json:"role" gorm:"column:role"`
	IDAuthor          int    `json:"id_author" gorm:"column:id_author"`
}

type UserEdit struct {
	Username    string      `json:"username" gorm:"column:username"`
	Email       string      `json:"email" gorm:"column:email"`
	Name        string      `json:"name" gorm:"column:name"`
	Telp        string      `json:"telp" gorm:"column:telp"`
	DateOfBirth interface{} `json:"dateOfBirth" gorm:"column:date_of_birth"`
}
type UserName struct {
	Username string `json:"username" gorm:"column:username"`
}

// ForgotPass is mode for request forgot pass
type ForgotPass struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}
