package controllers

import "gorm.io/gorm"

type DatabaseController struct {
	Database *gorm.DB
}
