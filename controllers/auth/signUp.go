package controllers

import (
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (db *DatabaseController) Signup(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string
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

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := db.Database.Create(&user)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
