package main

import (
	"be-cp2b/internal"
	"be-cp2b/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

// Swagger documentation
// @title BE CP2B - REST API Docs
// @description Simply cp2b system
// @version 1.0
// @BasePath /
// @schemes http
// @schemes https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	App := internal.InitApp()

	router.InitRouter(r, App)

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
