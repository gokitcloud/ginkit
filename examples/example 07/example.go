package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/gokitcloud/ginkit"
	"golang.org/x/oauth2"
)

var (
	authServerURL = "http://localhost:9096"
	oauthConfig   = oauth2.Config{
		ClientID:     "222222",
		ClientSecret: "22222222",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:3333/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}
)

func main() {
	e := ginkit.NewDefaultWithSessions("memstore", "example07", "123456578").SetVersion("0.0.0").AddHealthCheckFunc(MyInternalHealthCheck)

	restricted := e.OAuthGroup("/org/:id", authServerURL, oauthConfig)
	restricted.Use(ginkit.RBACMiddleware("org_model.conf", "org_policy.csv", "session:oauth_user_id", "param:id"))
	restricted.Use(ginkit.RemoveHeaders("X-Token"))
	restricted.Use(ginkit.RemoveHeaders("Cookie"))
	restricted.GET("", ginkit.WrapDataFuncParams(test2))
	restricted.GET("/", ginkit.WrapDataFuncParams(test2))
	restricted.GET("/proxy/*proxyPath", ginkit.Proxy("https://httpbin.org/"))

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
