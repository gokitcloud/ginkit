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
	restrictedA.GET("/:ping", func(p ginkit.Params) (any, error) {
		if ping, ok := p.Get("ping"); ok && ping == "ping" {
			return ginkit.H{
				"message": "pong",
			}, nil
		}

		return nil, errors.New("invalid ping")
	})

	r.Run(os.Getenv("PORT"))
}
