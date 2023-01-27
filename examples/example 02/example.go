package main

import (
	"log"

	"github.com/gokitcloud/ginkit"
)

func main() {
	e := ginkit.NewDefault()

	restricted := e.SimpleTokenAuthGroup("/", "12345678", "X-Token")
	restricted.GET("/test", ginkit.WrapDataFunc(test))

	err := e.Run(":3333")
	if err != nil {
		log.Println(err)
	}
}

func test() (any, error) {
	return map[string]any{"foo": "bar"}, nil
}
