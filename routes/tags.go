package routes

import (
	controllers "github.com/JohnKucharsky/jwt-golang/controllers/tag"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
)

func Tags(r *gin.Engine) {
	databaseController := controllers.DatabaseController{Database: initializers.DB}

	r.POST("/tags", middleware.RequireAuth, databaseController.CreateTag)
	r.GET("/tags", databaseController.GetTags)
	r.GET("/tags/:id", databaseController.GetOneTag)
	r.PUT("/tags/:id", middleware.RequireAuth, databaseController.UpdateTag)
	r.DELETE("/tags/:id", middleware.RequireAuth, databaseController.DeleteTag)
}
