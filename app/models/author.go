package models

// AuthorProfile is model for view author profile
type AuthorProfile struct {
	IDUser         int    `json:"id_user" gorm:"column:id"`
	IDAuthor       int    `json:"id_author" gorm:"column:id_author"`
	AuthorNickName string `json:"authorNickName" gorm:"column:author_nickname"`
	AuthorName     string `json:"authorName" gorm:"column:author_name"`
	AuthorAvatar   string `json:"authorAvatar" gorm:"column:author_avatar"`
	Bio            string `json:"bio" gorm:"column:bio"`
	Link           string `json:"link" gorm:"column:link"`
	Instagram      string `json:"instagram" gorm:"column:instagram"`
	Twitter        string `json:"twitter" gorm:"column:twitter"`
	Youtube        string `json:"youtube" gorm:"column:youtube"`
	Facebook       string `json:"facebook" gorm:"column:facebook"`
}

type AuthorNickName struct {
	AuthorNickName string `json:"authorNickName" gorm:"column:author_nickname"`
}

// AuthorData is model for table author profile Data
type AuthorData struct {
	AuthorID       string `json:"authorId" gorm:"column:id"`
	AuthorNickName string `json:"authorNickName" gorm:"column:nickname"`
	AuthorName     string `json:"authorName" gorm:"column:name"`
	AuthorAvatar   string `json:"authorAvatar" gorm:"column:avatar"`
	Bio            string `json:"bio" gorm:"column:bio"`
	Link           string `json:"link" gorm:"column:link"`
	Instagram      string `json:"instagram" gorm:"column:instagram"`
	Twitter        string `json:"twitter" gorm:"column:twitter"`
	Youtube        string `json:"youtube" gorm:"column:youtube"`
	Facebook       string `json:"facebook" gorm:"column:facebook"`
}

// AuthorConfig is model for config author profile Data
type AuthorConfig struct {
	AuthorID     string `json:"authorId" gorm:"column:id"`
	Deleted      string `json:"deleted" gorm:"column:deleted"`
	FlagApproval string `json:"flagApproval" gorm:"column:flag_approval"`
}
