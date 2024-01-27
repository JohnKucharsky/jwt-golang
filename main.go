package main

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/controllers"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	r.POST(
		"/signup", controllers.Signup,
	)
	r.POST(
		"/login", controllers.Login,
	)
	r.GET("/welcome", middleware.RequireAuth, controllers.Validate)
	r.POST("/refresh", controllers.RefreshToken)
	r.POST("/logout", controllers.Logout)

	err := r.Run(port)
	if err != nil {
		fmt.Println(err)
	}
}
