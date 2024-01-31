package controllers

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"strconv"
)

func (db *DatabaseController) CreateTag(c *gin.Context) {
	postId := c.Query("post")

	var body struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color"`
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

	tag := models.Tag{Name: body.Name, Color: body.Color}

	result := db.Database.Model(&models.Tag{}).Preload("Posts").Create(&tag)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	// append post
	var posts []models.Post
	err = db.Database.Model(&tag).Association("Posts").Find(&posts)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	parsedId, err := strconv.ParseUint(postId, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	_, ok := lo.Find(
		posts, func(item models.Post) bool {
			return item.ID == uint(parsedId)
		},
	)

	if postId != "" && !ok {
		var post models.Post
		postResult := db.Database.First(&post, postId)

		if postResult.Error != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{
					"error": postResult.Error.Error(),
				},
			)
			return
		}

		posts = append(posts, post)
	}

	err = db.Database.Model(&tag).Association("Posts").Append(
		posts,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	// append post end

	c.JSON(
		http.StatusOK, gin.H{
			"data": tag,
		},
	)
}
