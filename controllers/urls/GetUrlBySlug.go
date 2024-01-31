package controllers

import (
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (db *DatabaseController) GetUrlBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var url models.ShortUrl
	err := db.Database.Where("Slug = ?", slug).First(&url).Error

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"data": url,
		},
	)
}
