package routes

import (
	"github.com/JohnKucharsky/jwt-golang/controllers"
	"github.com/JohnKucharsky/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
)

func Tags(r *gin.Engine) {
	r.POST("/tags", middleware.RequireAuth, controllers.CreateTag)
	r.GET("/tags", controllers.GetTags)
	r.GET("/tags/:id", controllers.GetOneTag)
	r.PUT("/tags/:id", middleware.RequireAuth, controllers.UpdateTag)
	r.DELETE("/tags/:id", middleware.RequireAuth, controllers.DeleteTag)
}
