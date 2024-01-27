package controllers

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTag(c *gin.Context) {
	var body struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color"`
	}

	err := c.Bind(&body)
	fmt.Println(err)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	tag := models.Tag{Name: body.Name, Color: body.Color}

	result := initializers.DB.Create(&tag)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
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

func GetTags(c *gin.Context) {
	var tags []models.Tag
	initializers.DB.Find(&tags)

	c.JSON(
		http.StatusOK, gin.H{
			"data": tags,
		},
	)
}

func GetOneTag(c *gin.Context) {
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

	c.JSON(
		http.StatusOK, gin.H{
			"data": tag,
		},
	)
}

func UpdateTag(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color"`
	}

	err := c.Bind(&body)
	fmt.Println(err)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

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

	result = initializers.DB.Model(&tag).Updates(
		models.Tag{
			Name:  body.Name,
			Color: body.Color,
		},
	)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
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
