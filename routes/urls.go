package routes

import (
	controllers "github.com/JohnKucharsky/jwt-golang/controllers/urls"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
)

func Urls(r *gin.Engine) {
	databaseController := controllers.DatabaseController{Database: initializers.DB}

	r.POST("/urls", middleware.RequireAuth, databaseController.CreateShortUrl)
	r.GET("/urls/:slug", databaseController.GetUrlBySlug)

}
