package models

// User is model for user data
type User struct {
	ID                int         `json:"id_user" gorm:"column:id"`
	Username          string      `json:"username" gorm:"column:username"`
	Email             string      `json:"email" gorm:"column:email"`
	Password          string      `json:"password" gorm:"column:password"`
	Name              string      `json:"name" gorm:"column:name"`
	Telp              string      `json:"telp" gorm:"column:telp"`
	DateOfBirth       interface{} `json:"dateOfBirth" gorm:"column:date_of_birth"`
	Google            int         `json:"google" gorm:"column:google"`
	IDRole            int         `json:"idRole" gorm:"column:id_role"`
	EmailVerification int         `json:"emailVerify" gorm:"column:email_verify"`
	User              string      `json:"user" gorm:"-"`
}

// ForgotPass is mode for request forgot pass
type ForgotPass struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}
