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
	r.Run(os.Getenv("PORT"))
}
