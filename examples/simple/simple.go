package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gokitcloud/ginkit"
)

func main() {
	r := ginkit.Default()
	r.GET("/ping", gin.H{
		"message": "pong",
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(os.Getenv(port))
}
