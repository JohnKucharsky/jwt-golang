package controllers

import (
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	err := initializers.DB.Model(&models.Post{}).Preload("Tags").Find(&posts).Error

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
			"data": posts,
		},
	)
}

func GetOnePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	err := initializers.DB.Model(&models.Post{}).Preload("Tags").First(
		&post,
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
			"data": post,
		},
	)
}
