package main

import (
	"errors"

	"github.com/gokitcloud/ginkit"
)

func main() {

	r := ginkit.Default()
	restricted := r.SimpleTokenAuthGroup("/", "12345678", "X-Token")
	restricted.GET("/:ping", func(p ginkit.Params) (any, error) {
		if ping, ok := p.Get("ping"); ok && ping == "ping" {
			return ginkit.H{
				"message": "pong",
			}, nil
		}

		return nil, errors.New("invalid ping")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
