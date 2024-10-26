package main

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gokitcloud/ginkit"
)

func main() {
	r := ginkit.Default()
	restricted := r.SimpleTokenAuthGroup("/", "12345678", "X-Token")
	restricted.GET("/ping", gin.H{
		"message": "pong",
	})

	restrictedA := r.SimpleTokenAuthGroup("/a", "12345678", "X-Token")
	restrictedA.GET("/:path", func(p ginkit.Params) (any, error) {
		if ping, ok := p.Get("path"); ok && ping == "ping" {
			return ginkit.H{
				"message": "pong",
			}, nil
		}

		return nil, errors.New("invalid ping")
	})

	port := os.Getenv("SA_PORT")
	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
