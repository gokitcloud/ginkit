package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/gokitcloud/ginkit"
)

func main() {
	e := ginkit.NewDefault().SetVersion("0.0.0").AddHealthCheckFunc(MyInternalHealthCheck)

	restricted := e.RBACTokenPathGroup("/", "path_model.conf", "path_policy.csv", "X-Token")
	restricted.GET("/test", ginkit.WrapDataFunc(test))

	restricted2 := e.RBACGroup("/org/:id", "org_model.conf", "org_policy.csv", "header:X-Token", "param:id")
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

func MyInternalHealthCheck() error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return errors.New("Random error")
	}
	return nil
}
