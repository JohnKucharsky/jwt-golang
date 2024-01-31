package controllers

import (
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (db *DatabaseController) CreateShortUrl(c *gin.Context) {
	var body struct {
		Destination string `json:"destination" binding:"required"`
		Slug        string `json:"slug"`
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	url := models.ShortUrl{Destination: body.Destination, Slug: body.Slug}

	err = db.Database.Create(&url).Error

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
