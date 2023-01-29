package main

import (
	"github.com/gokitcloud/ginkit"
)

func main() {
	r := ginkit.Default()
	restricted := r.SimpleTokenAuthGroup("/", "12345678", "X-Token")
	restricted.GET("/ping", ginkit.H{
		"message": "pong",
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
