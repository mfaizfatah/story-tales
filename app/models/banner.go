package models

import (
	"mime/multipart"
	"time"
)

// Story model data
type BannerReq struct {
	Category     string                `json:"category"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	URL          string                `json:"url"`
	Status       int                   `json:"status"`
	DaysValid    int                   `json:"daysValid"`
	Sequence     int                   `json:"sequence"`
	DeepLink     bool                  `json:"deepLink"`
	DetailStatus bool                  `json:"detailStatus"`
	ServiceID    string                `json:"serviceId"`
	ImgFile      *multipart.FileHeader `form:"imgFile" binding:"required"`
	ThumbFile    *multipart.FileHeader `form:"thumbFile" binding:"required"`
}

type BannerRequ struct {
	Id        int                   `uri:"id"`
	Category  string                `form:"category"`
	Title     string                `form:"title"`
	URL       string                `form:"url"`
	ImgFile   *multipart.FileHeader `form:"imgFile" binding:"required"`
	ThumbFile *multipart.FileHeader `form:"thumbFile" binding:"required"`
}

type Banner struct {
	ID           int       `json:"id"`
	Category     string    `json:"category"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	URL          string    `json:"url"`
	Image        string    `json:"image"`
	Thumb        string    `json:"thumb"`
	Status       int       `json:"status"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
	ValidUntil   time.Time `json:"validUntil" gorm:"column:validUntil"`
	Sequence     int       `json:"sequence"`
	DeepLink     bool      `json:"deepLink" gorm:"column:deepLink"`
	DetailStatus bool      `json:"detailStatus" gorm:"column:detailStatus"`
	ServiceID    string    `json:"serviceId" gorm:"column:serviceId"`
}

type BannerRs struct {
	Id             int       `json:"id"`
	Category       string    `json:"category"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Url            string    `json:"url"`
	Image          string    `json:"image"`
	Thumb          string    `json:"thumb"`
	Status         int       `json:"status"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at"`
	ValidUntil     time.Time `json:"validUntil" gorm:"column:validUntil"`
	Sequence       int       `json:"sequence"`
	DeepLink       bool      `json:"deepLink" gorm:"column:deepLink"`
	DetailStatus   bool      `json:"detailStatus" gorm:"column:detailStatus"`
	ServiceId      string    `json:"serviceId" gorm:"column:serviceId"`
	ResponseStatus string    `json:"responseStatus" gorm:"column:responseStatus"`
}

type BannerDetailRs struct {
	Id           int       `json:"id"`
	Category     string    `json:"category"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Url          string    `json:"url"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at"`
	ValidUntil   time.Time `json:"validUntil" gorm:"column:validUntil"`
	DeepLink     bool      `json:"deepLink" gorm:"column:deepLink"`
	DetailStatus bool      `json:"detailStatus" gorm:"column:detailStatus"`
	ServiceId    string    `json:"serviceId" gorm:"column:serviceId"`
}

// Response All..
type ListBannerThumbRs struct {
	Id           int    `json:"id"`
	Sequence     int    `json:"sequence"`
	Category     string `json:"category"`
	Title        string `json:"title"`
	Thumb        string `json:"thumb"`
	DeepLink     bool   `json:"deepLink" gorm:"column:deepLink"`
	DetailStatus bool   `json:"detailStatus" gorm:"column:detailStatus"`
	ServiceId    string `json:"serviceId" gorm:"column:serviceId"`
}
