package controllers

import (
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (db *DatabaseController) DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := db.Database.First(&post, id)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	result = db.Database.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	c.Status(http.StatusOK)
}
