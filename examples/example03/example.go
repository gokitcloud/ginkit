package main

import (
	"errors"
	"log"

	"github.com/gokitcloud/ginkit"
)

func main() {
	e := ginkit.NewDefault()

	restricted := e.SimpleTokenAuthGroup("/", "12345678", "X-Token")
	restricted.GET("/test", ginkit.WrapDataFunc(test))
	restricted.GET("/test/:id/*path", ginkit.WrapDataFuncParams(test2))

	err := e.Run(":3333")
	if err != nil {
		log.Println(err)
	}
}

func test() (any, error) {
	return map[string]any{"foo": "bar"}, nil
}

func test2(p ginkit.Params) (any, error) {
	if id, _ := p.Get("id"); id != "123" {
		return nil, errors.New("invalid id")
	}
	return map[string]any{
		"foo":    "bar",
		"params": p,
	}, nil
}
