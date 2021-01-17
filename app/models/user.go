package models

// User is model for user data
type User struct {
	ID          int         `json:"id_user" gorm:"column:id"`
	Username    string      `json:"username" gorm:"column:username"`
	Email       string      `json:"email" gorm:"column:email"`
	Password    string      `json:"password" gorm:"column:password"`
	Name        string      `json:"name" gorm:"column:name"`
	Telp        string      `json:"telp" gorm:"column:telp"`
	DateOfBirth interface{} `json:"dateOfBirth" gorm:"column:date_of_birth"`
	IDRole      int         `json:"idRole" gorm:"column:id_role"`
}
