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

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	tagId := c.Query("tag")

	var body struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body"`
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

	// append tag
	var tags []models.Tag
	err = initializers.DB.Model(&post).Association("Tags").Find(&tags)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	parsedId, err := strconv.ParseUint(tagId, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	_, ok := lo.Find(
		tags, func(item models.Tag) bool {
			return item.ID == uint(parsedId)
		},
	)

	if tagId != "" && !ok {
		var tag models.Tag
		tagResult := initializers.DB.First(&tag, tagId)

		if tagResult.Error != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{
					"error": tagResult.Error.Error(),
				},
			)
			return
		}

		tags = append(tags, tag)
	}

	err = initializers.DB.Model(&post).Association("Tags").Append(
		tags,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	// append tag end

	c.JSON(
		http.StatusOK, gin.H{
			"data": post,
		},
	)
}
