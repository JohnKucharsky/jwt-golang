package controllers

import (
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(c *gin.Context) {
	var tags []models.Tag
	err := initializers.DB.Model(&models.Tag{}).Preload("Posts").Find(&tags).Error

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
			"data": tags,
		},
	)
}

func GetOneTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	err := initializers.DB.Model(&models.Tag{}).Preload("Posts").First(
		&tag,
		id,
	).Error

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
			"data": tag,
		},
	)
}
