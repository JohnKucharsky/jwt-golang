package models

import "gorm.io/gorm"

type ShortUrl struct {
	gorm.Model
	Destination string `gorm:"unique;not null" json:"destination"`
	Slug        string `json:"slug"`
}
