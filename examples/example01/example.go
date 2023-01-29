package main

import (
	"log"

	"github.com/gokitcloud/ginkit"
)

func main() {
	e := ginkit.NewDefault()

	e.GET("/test", test)

	err := e.Run(":3333")
	if err != nil {
		log.Println(err)
	}
}

func test() (any, error) {
	return map[string]any{"foo": "bar"}, nil
}
