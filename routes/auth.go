package routes

import (
	"github.com/JohnKucharsky/jwt-golang/controllers/auth"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	databaseController := controllers.DatabaseController{Database: initializers.DB}

	r.POST(
		"/signup", databaseController.Signup,
	)
	r.POST(
		"/login", databaseController.Login,
	)
	r.POST("/refresh", controllers.RefreshToken)
	r.POST("/logout", controllers.Logout)
}
