package main

import (
	"be-cp2b/internal"
	"be-cp2b/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

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
