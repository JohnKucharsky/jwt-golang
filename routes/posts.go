package routes

import (
	controllers "github.com/JohnKucharsky/jwt-golang/controllers/post"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
)

func Posts(r *gin.Engine) {
	databaseController := controllers.DatabaseController{Database: initializers.DB}

	r.POST("/posts", middleware.RequireAuth, databaseController.CreatePost)
	r.GET("/posts", databaseController.GetPosts)
	r.GET("/posts/:id", databaseController.GetOnePost)
	r.PUT("/posts/:id", middleware.RequireAuth, databaseController.UpdatePost)
	r.DELETE(
		"/posts/:id",
		middleware.RequireAuth,
		databaseController.DeletePost,
	)
}
