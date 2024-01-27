package controllers

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body"`
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

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

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
			"data": post,
		},
	)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(
		http.StatusOK, gin.H{
			"data": posts,
		},
	)
}

func GetOnePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

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
			"data": post,
		},
	)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body"`
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

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	result = initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
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
			"data": post,
		},
	)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	result = initializers.DB.Delete(&models.Post{}, id)

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
