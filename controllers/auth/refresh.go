package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func RefreshToken(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(
			http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"Unexpected signing method: %v",
					token.Header["alg"],
				)
			}

			return []byte(os.Getenv("SECRET")), nil

		},
	)

	if err != nil {
		c.JSON(
			http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error": "Expired",
				},
			)
			return
		}

		expirationTime := time.Now().Add(10 * time.Minute)
		claims["exp"] = jwt.NewNumericDate(expirationTime)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		c.SetCookie("Authorization", tokenString, 60*5, "", "", false, true)
	} else {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Can't refresh token",
			},
		)
		return
	}

}
