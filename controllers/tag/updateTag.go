package controllers

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"strconv"
)

func UpdateTag(c *gin.Context) {
	id := c.Param("id")
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

	// append post
	var posts []models.Post
	err = initializers.DB.Model(&tag).Association("Posts").Find(&posts)

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
		postResult := initializers.DB.First(&post, postId)

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

	err = initializers.DB.Model(&tag).Association("Posts").Append(
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
