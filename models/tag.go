package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string  `gorm:"unique;not null" json:"name"`
	Color string  `json:"color"`
	Posts []*Post `gorm:"many2many:posts_tags" json:"posts"`
}
