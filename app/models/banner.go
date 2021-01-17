package models

import (
	"time"
)

// Story model data
type Banner struct {
	ID           int       `json:"id"`
	Category     string    `json:"category"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	URL          string    `json:"url"`
	Image        string    `json:"image"`
	Thumb        string    `json:"thumb"`
	Status       int       `json:"status"`
	CreateAt     time.Time `json:"createAt" gorm:"column:createAt"`
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
	CreateAt       time.Time `json:"createAt" gorm:"column:createAt"`
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
	CreateAt     time.Time `json:"createAt" gorm:"column:createAt"`
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
