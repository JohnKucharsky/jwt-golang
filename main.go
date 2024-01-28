package main

import (
	"fmt"
	"github.com/JohnKucharsky/jwt-golang/initializers"
	"github.com/JohnKucharsky/jwt-golang/routes"
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
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Println(err)
	}

	routes.Auth(r)
	routes.Posts(r)
	routes.Tags(r)

	err = r.Run(port)
	if err != nil {
		fmt.Println(err)
	}
}
