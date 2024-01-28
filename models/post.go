package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"unique;not null" json:"title"`
	Body  string `json:"body"`
	Tags  []*Tag `gorm:"many2many:posts_tags" json:"tags"`
}
