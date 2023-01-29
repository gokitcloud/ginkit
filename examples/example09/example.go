package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/gokitcloud/ginkit"
)

var (
	samlConfig   = &ginkit.SAMLGroupConfig{
		MetaDataFile: "metadata.xml",
		CertFile:     "myservice.cert",
		KeyFile:      "myservice.key",
		EntityID:     "asdfasdf",
		RootURL:      "https://devlocal.site:8000",
		ParamMap: map[string]string{
			"email": "email",
		},
	}
	rbacConfig = &ginkit.RBACConfig{
		Model:  "org_model.conf",
		Policy: "org_policy.csv",
		Params: []any{
			"param:email",
			"param:id",
		},
	}
)

func main() {
	e := ginkit.NewDefaultWithSessions("memstore", "example07", "123456578").SetVersion("0.0.0").AddHealthCheckFunc(MyInternalHealthCheck)

	restricted := e.SAMLGroup("/org/:id", samlConfig)
	restricted.Use(ginkit.RBACMiddleware(rbacConfig))
	restricted.Use(ginkit.RemoveHeaders("Cookie"))
	restricted.Use(ginkit.AddRequestHeader("email", "email"))
	restricted.GET("", ginkit.WrapDataFuncParams(test2))
	restricted.GET("/", ginkit.WrapDataFuncParams(test2))
	restricted.GET("/proxy/*proxyPath", ginkit.Proxy("https://httpbin.org/"))

	err := e.RunSSLSelfSigned(":8000")
	if err != nil {
		log.Println(err)
	}
}

func test() (any, error) {
	return map[string]any{"foo": "bar"}, nil
}

func test2(p ginkit.Params) (any, error) {
	id, _ := p.Get("id")
	email, _ := p.Get("email")

	return map[string]any{
		"foo":    "bar",
		"orgid":  id,
		"email":  email,
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
