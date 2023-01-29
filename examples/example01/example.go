package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gokitcloud/ginkit"
)

func main() {
	r := ginkit.Default()
	r.GET("/ping", gin.H{
		"message": "pong",
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
