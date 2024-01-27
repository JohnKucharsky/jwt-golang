package routes

import (
	"github.com/JohnKucharsky/jwt-golang/controllers"
	"github.com/JohnKucharsky/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	r.POST(
		"/signup", controllers.Signup,
	)
	r.POST(
		"/login", controllers.Login,
	)
	r.GET("/welcome", middleware.RequireAuth, controllers.Validate)
	r.POST("/refresh", controllers.RefreshToken)
	r.POST("/logout", controllers.Logout)
}
