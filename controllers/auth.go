package controllers

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			},
		)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Failed to hash password",
			},
		)
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Failed to create user",
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			},
		)
		return
	}

	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			},
		)

		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(body.Password),
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			},
		)
		return
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Failed to create token",
			},
		)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(
		http.StatusOK, gin.H{"data": user},
	)
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(
		http.StatusOK, gin.H{
			"message": user,
		},
	)
}

func RefreshToken(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"Unexpected signing method: %v",
					token.Header["alg"],
				)
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil

		},
	)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(10 * time.Minute)
		claims["exp"] = jwt.NewNumericDate(expirationTime)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{
					"error": "Failed to create token",
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

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
}
