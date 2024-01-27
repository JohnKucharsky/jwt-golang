package routes

import (
	"github.com/JohnKucharsky/jwt-golang/controllers"
	"github.com/JohnKucharsky/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
)

func Posts(r *gin.Engine) {
	r.POST("/posts", middleware.RequireAuth, controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetOnePost)
	r.PUT("/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	r.DELETE("/posts/:id", middleware.RequireAuth, controllers.DeletePost)
}
