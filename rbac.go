package ginkit

import (
	"log"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func (e *Engine) RBACTokenPathGroup(path string, model, policy any, tokenHeader string) *gin.RouterGroup {
	restricted := e.Router().Group(path)
	restricted.Use(RBACTokenPathMiddleware(model, policy, tokenHeader))

	return restricted
}

func RBACTokenPathMiddleware(model, policy any, tokenHeader string) func(c *gin.Context) {
	e, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		allowed := false
		method := c.Request.Method
		path := c.Request.URL.Path

		for _, reqToken := range c.Request.Header[tokenHeader] {
			allowed, err = e.Enforce(reqToken, path, method)
			if err != nil {
				ReturnData(c, http.StatusForbidden, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			if allowed {
				continue
			}
		}

		if !allowed {
			ReturnData(c, http.StatusForbidden, gin.H{"error": "RBAC did not allow authorization"})
			c.Abort()
			return
		}
	}
}

func (e *Engine) RBACTokenParamGroup(path string, model, policy any, tokenHeader string, params []string) *gin.RouterGroup {
	restricted := e.Router().Group(path)
	restricted.Use(RBACTokenParamMiddleware(model, policy, tokenHeader, params))

	return restricted
}

func RBACTokenParamMiddleware(model, policy any, tokenHeader string, params []string) func(c *gin.Context) {
	e, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		allowed := false
		var paramValues []string

		for _, p := range params {
			paramValues = append(paramValues, c.Param(p))
		}

		for _, reqToken := range c.Request.Header[tokenHeader] {
			args := []interface{}{reqToken}
			for _, v := range paramValues {
				args = append(args, v)
			}
			allowed, err = e.Enforce(args...)
			if err != nil {
				ReturnData(c, http.StatusForbidden, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			if allowed {
				continue
			}
		}

		if !allowed {
			ReturnData(c, http.StatusForbidden, gin.H{"error": "RBAC did not allow authorization"})
			c.Abort()
			return
		}
	}
}

func (e *Engine) RBACGroup(path string, model, policy any, params ...any) *gin.RouterGroup {
	restricted := e.Router().Group(path)
	restricted.Use(RBACMiddleware(model, policy, params...))

	return restricted
}

func RBACMiddleware(model, policy any, params ...any) func(c *gin.Context) {
	e, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		allowed := false
		var paramValues []any

		for _, p := range params {
			var paramValue any

			switch p.(type) {
			case string:
				param := strings.Split(p.(string), ":")
				if len(param) < 2 {
					ReturnData(c, http.StatusForbidden, gin.H{"error": "invalid RBAC config"})
					c.Abort()
					return
				}

				switch param[0] {
				case "request":
					switch param[1] {
					case "method":
						paramValue = c.Request.Method
					case "path":
						paramValue = c.Request.URL.Path
					}
				case "value":
					paramValue = param[1]
				case "param":
					paramValue = c.Param(param[1])
				case "header":
					paramValue = c.Request.Header.Get(param[1])
				case "session":
				}
			case func(c *gin.Context) any:
				paramValue = p.(func(c *gin.Context) any)(c)
			}

			if paramValue == nil {
				ReturnData(c, http.StatusForbidden, gin.H{"error": "RBAC did not allow authorization"})
				c.Abort()
				return
			}

			paramValues = append(paramValues, paramValue)
		}

		allowed, err = e.Enforce(paramValues...)
		if err != nil {
			log.Println(err)
			log.Println(paramValues)
			ReturnData(c, http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !allowed {
			ReturnData(c, http.StatusForbidden, gin.H{"error": "RBAC did not allow authorization"})
			c.Abort()
			return
		}
	}
}
