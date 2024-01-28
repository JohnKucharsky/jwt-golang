package controllers

import (
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	result := initializers.DB.First(&tag, id)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	result = initializers.DB.Delete(&models.Tag{}, id)

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
