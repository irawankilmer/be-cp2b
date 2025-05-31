package main

import (
	"be-cp2b/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	internal.InitApp()

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
