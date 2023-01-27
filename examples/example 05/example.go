package main

import (
	"log"

	"github.com/gokitcloud/ginkit"
)

func main() {
	e := ginkit.NewDefault().SetVersion("0.0.0")

	restricted := e.RBACPathGroup("/", "path_model.conf", "path_policy.csv", "X-Token")
	restricted.GET("/test", ginkit.WrapDataFunc(test))

	restricted2 := e.RBACParamGroup("/org/:id", "org_model.conf", "org_policy.csv", "X-Token", []string{"id"})
	restricted2.GET("", ginkit.WrapDataFunc(test))
	restricted2.GET("/", ginkit.WrapDataFunc(test))

	err := e.Run(":3333")
	if err != nil {
		log.Println(err)
	}
}

func test() (any, error) {
	return map[string]any{"foo": "bar"}, nil
}

func test2(p ginkit.Params) (any, error) {
	id, _ := p.Get("id")

	return map[string]any{
		"foo":    "bar",
		"orgid":  id,
		"params": p,
	}, nil
}
